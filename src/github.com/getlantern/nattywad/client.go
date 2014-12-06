package nattywad

import (
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/getlantern/go-natty/natty"
	"github.com/getlantern/waddell"
)

// ServerPeer identifies a server for NAT traversal
type ServerPeer struct {
	// ID: the server's PeerID on waddell (type 4 GUID)
	ID string

	// WaddellAddr: the address of the waddell server on which the server is
	// listening for offers.
	WaddellAddr string

	// Extras: Extra information about the peer (pass-through)
	Extras map[string]interface{}
}

func (p *ServerPeer) CompositeID() string {
	return p.WaddellAddr + "|" + p.ID
}

func (p *ServerPeer) String() string {
	return p.CompositeID()
}

// SuccessCallbackClient is a function that gets invoked when a client NAT
// traversal results in a UDP five tuple.
type SuccessCallbackClient func(info *TraversalInfo)

// FailureCallbackClient is a callback that is invoked if a client NAT traversal
// fails.
type FailureCallbackClient func(info *TraversalInfo)

// TraversalInfo provides information about failed traversals
type TraversalInfo struct {
	Peer *ServerPeer
	// ServerRespondedToSignaling: indicates whether nattywad received any
	// signaling messages from the server peer during the traversal.
	ServerRespondedToSignaling bool

	// ServerGotFiveTuple: indicates whether or not the server peer got a
	// FiveTuple of its own (only populated when using nattywad as a client)
	ServerGotFiveTuple bool

	// LocalAddr: on a successful traversal, this contains the local UDP addr of
	// the FiveTuple.
	LocalAddr *net.UDPAddr

	// RemoteAddr: on a successful traversal, this contains the remote UDP addr
	// of the FiveTuple.
	RemoteAddr *net.UDPAddr

	// Duration: the duration of the traversal
	Duration time.Duration
}

// Client is a client that initiates NAT traversals to one or more configured
// servers. When a NAT traversal results in a 5-tuple, the OnFiveTuple callback
// is called.
type Client struct {
	// DialWaddell: a function that dials the waddell server at the given
	// address. Must be specified in order for Client to work.
	DialWaddell func(addr string) (net.Conn, error)

	// ServerCert: PEM-encoded certificate by which to authenticate the waddell
	// server. If provided, connection to waddell is encrypted with TLS. If not,
	// connection will be made plain-text.
	ServerCert string

	// OnSuccess: a callback that's invoked once a five tuple has been
	// obtained. Must be specified in order for Client to work.
	OnSuccess SuccessCallbackClient

	// OnFailure: a callback that's invoked if the NAT traversal fails
	// (e.g. times out). If unpopulated, failures aren't reported.
	OnFailure FailureCallbackClient

	// KeepAliveInterval: If specified to a non-zero value, nattywad will send a
	// keepalive message over the waddell channel to keep open the underlying
	// connections.
	KeepAliveInterval time.Duration

	serverPeers  map[string]*ServerPeer
	workers      map[uint32]*clientWorker
	workersMutex sync.RWMutex
	waddellConns map[string]*waddellConn
	cfgMutex     sync.Mutex
}

// Configure (re)configures this Client to communicate with the given list of
// server peers. Anytime that the list is found to contain a new peer, a NAT
// traversal is attempted to that peer.
func (c *Client) Configure(serverPeers []*ServerPeer) {
	c.cfgMutex.Lock()
	defer c.cfgMutex.Unlock()

	log.Debugf("Configuring nat traversal client with %d server peers", len(serverPeers))

	// Lazily initialize data structures
	if c.serverPeers == nil {
		c.serverPeers = make(map[string]*ServerPeer)
		c.waddellConns = make(map[string]*waddellConn)
		c.workers = make(map[uint32]*clientWorker)
	}

	priorServerPeers := c.serverPeers
	c.serverPeers = make(map[string]*ServerPeer)

	for _, peer := range serverPeers {
		cid := peer.CompositeID()

		if priorServerPeers[cid] == nil {
			// Either we have a new server, or the address changed, try to
			// traverse
			log.Debugf("Attempting traversal to %s", peer.ID)
			peerId, err := waddell.PeerIdFromString(peer.ID)
			if err != nil {
				log.Errorf("Unable to parse PeerID for server peer %s: %s",
					peer.ID, err)
				continue
			}
			c.offer(peer, peerId)
		} else {
			log.Debugf("Already know about %s, not attempting traversal", peer.ID)
		}

		// Keep track of new peer
		c.serverPeers[cid] = peer
	}
}

func (c *Client) offer(serverPeer *ServerPeer, peerId waddell.PeerId) {
	wc := c.waddellConns[serverPeer.WaddellAddr]
	if wc == nil {
		/* new waddell server--open connection to it */
		var err error
		wc, err = newWaddellConn(func() (net.Conn, error) {
			return c.DialWaddell(serverPeer.WaddellAddr)
		}, c.ServerCert)
		if err != nil {
			log.Errorf("Unable to connect to waddell: %s", err)
			return
		}
		if c.KeepAliveInterval > 0 {
			// Periodically send a KeepAlive message
			go func() {
				for {
					time.Sleep(c.KeepAliveInterval)
					err := wc.client.SendKeepAlive()
					if err != nil {
						log.Errorf("Unable to send KeepAlive packet to waddell: %s", err)
						return
					}
				}
			}()
		}
		c.waddellConns[serverPeer.WaddellAddr] = wc
		go c.receiveMessages(wc)
	}

	w := &clientWorker{
		wc:          wc,
		peerId:      peerId,
		onSuccess:   c.OnSuccess,
		onFailure:   c.OnFailure,
		traversalId: uint32(rand.Int31()),
		info:        &TraversalInfo{Peer: serverPeer},
		serverReady: make(chan bool, 10), // make this buffered to prevent deadlocks
	}
	c.addWorker(w)
	go w.run()
}

func (c *Client) receiveMessages(wc *waddellConn) {
	for {
		msg, _, err := wc.receive()
		if err != nil {
			log.Errorf("Unable to receive next message from waddell: %s", err)
			return
		}
		w := c.getWorker(msg.getTraversalId())
		if w == nil {
			log.Debugf("Got message for unknown traversal %d, skipping", msg.getTraversalId())
			continue
		}
		w.messageReceived(msg)
	}
}

func (c *Client) addWorker(w *clientWorker) {
	c.workersMutex.Lock()
	defer c.workersMutex.Unlock()
	c.workers[w.traversalId] = w
}

func (c *Client) getWorker(traversalId uint32) *clientWorker {
	c.workersMutex.RLock()
	defer c.workersMutex.RUnlock()
	return c.workers[traversalId]
}

func (c *Client) removeWorker(w *clientWorker) {
	c.workersMutex.Lock()
	defer c.workersMutex.Unlock()
	delete(c.workers, w.traversalId)
}

// clientWorker encapsulates the work done by the client for a single NAT
// traversal.
type clientWorker struct {
	wc          *waddellConn
	peerId      waddell.PeerId
	onSuccess   SuccessCallbackClient
	onFailure   FailureCallbackClient
	traversalId uint32
	traversal   *natty.Traversal
	info        *TraversalInfo
	startedAt   time.Time
	serverReady chan bool
}

func (w *clientWorker) run() {
	w.traversal = natty.Offer()
	defer w.traversal.Close()

	go w.sendMessages()

	w.startedAt = time.Now()
	ft, err := w.traversal.FiveTupleTimeout(Timeout)
	if err != nil {
		log.Errorf("Traversal to %s failed: %s", w.peerId, err)
		if w.onFailure != nil {
			w.info.Duration = time.Now().Sub(w.startedAt)
			w.onFailure(w.info)
		}
		return
	}
	if <-w.serverReady {
		local, remote, err := ft.UDPAddrs()
		if err != nil {
			log.Errorf("Unable to get UDP addresses for FiveTuple: %s", err)
			return
		}
		w.info.LocalAddr = local
		w.info.RemoteAddr = remote
		w.info.Duration = time.Now().Sub(w.startedAt)
		w.onSuccess(w.info)
	}
}

func (w *clientWorker) sendMessages() {
	for {
		msgOut, done := w.traversal.NextMsgOut()
		if done {
			return
		}
		w.wc.send(w.peerId, w.traversalId, msgOut)
	}
}

func (w *clientWorker) messageReceived(msg message) {
	msgString := string(msg.getData())

	// Update info
	w.info.ServerRespondedToSignaling = true
	if natty.IsFiveTuple(msgString) {
		w.info.ServerGotFiveTuple = true
	}

	if msgString == ServerReady {
		// Server's ready!
		w.serverReady <- true
	} else {
		w.traversal.MsgIn(msgString)
	}
}

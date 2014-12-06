// package nattywad implements NAT traversal using go-natty and waddell.
package nattywad

import (
	"encoding/binary"
	"net"
	"sync"
	"time"

	"github.com/getlantern/golog"
	"github.com/getlantern/waddell"
)

const (
	ServerReady = "ServerReady"
	Timeout     = 30 * time.Second
)

var (
	log = golog.LoggerFor("nattywad")

	maxWaddellMessageSize = 4096 + waddell.WADDELL_OVERHEAD

	endianness = binary.LittleEndian
)

// ConnectCallback is a function that gets invoked whenever a connection has
// been established to waddell.
type ConnectCallback func(id waddell.PeerId)

type message []byte

func (msg message) setTraversalId(id uint32) {
	endianness.PutUint32(msg[:4], id)
}

func (msg message) getTraversalId() uint32 {
	return endianness.Uint32(msg[:4])
}

func (msg message) getData() []byte {
	return msg[4:]
}

// waddellConn represents a connection to waddell. It automatically redials if
// there is a problem reading or writing.  As implemented, it provides
// half-duplex communication even though waddell supports full-duplex.  Given
// the low bandwidth and weak latency requirements of signaling traffic, this is
// unlikely to be a problem.
type waddellConn struct {
	dial       func() (net.Conn, error)
	serverCert string
	client     *waddell.Client
	conn       net.Conn
	connMutex  sync.RWMutex
	dialErr    error
}

// newWaddellConn establishes a new waddellConn that uses the provided dial
// function to connect to waddell when it needs to. serverCert specifies a PEM-
// encoded certificate with which to authenticate the waddell server.
func newWaddellConn(dial func() (net.Conn, error), serverCert string) (wc *waddellConn, err error) {
	wc = &waddellConn{
		dial:       dial,
		serverCert: serverCert,
	}
	err = wc.connect()
	return
}

func (wc *waddellConn) send(peerId waddell.PeerId, sessionId uint32, msgOut string) (err error) {
	log.Tracef("Sending message %s to peer %s in session %d", msgOut, peerId, sessionId)
	client, dialErr := wc.getClient()
	if dialErr != nil {
		err = dialErr
		return
	}
	err = client.SendPieces(peerId, idToBytes(sessionId), []byte(msgOut))
	if err != nil {
		wc.connError(client, err)
	}
	return
}

func (wc *waddellConn) receive() (msg message, from waddell.PeerId, err error) {
	log.Trace("Receiving")
	client, dialErr := wc.getClient()
	if dialErr != nil {
		err = dialErr
		return
	}
	b := make([]byte, maxWaddellMessageSize)
	var wm *waddell.Message
	wm, err = wc.client.Receive(b)
	if err != nil {
		wc.connError(client, err)
		log.Tracef("Error receiving: %s", err)
		return
	}
	from = wm.From
	msg = message(wm.Body)
	log.Tracef("Received %s from %s", msg.getData(), from)
	return
}

func (wc *waddellConn) getClient() (*waddell.Client, error) {
	wc.connMutex.RLock()
	defer wc.connMutex.RUnlock()
	return wc.client, wc.dialErr
}

func (wc *waddellConn) connError(client *waddell.Client, err error) {
	wc.connMutex.Lock()
	log.Tracef("Error on waddell connection: %s", err)
	if client == wc.client {
		// The current client is in error, redial
		go func() {
			log.Tracef("Redialing waddell")
			wc.dialErr = wc.connect()
			wc.connMutex.Unlock()
		}()
	} else {
		wc.connMutex.Unlock()
	}
}

func (wc *waddellConn) connect() (err error) {
	log.Trace("Connecting to waddell")
	wc.conn, err = wc.dial()
	if err != nil {
		return
	}
	if wc.serverCert != "" {
		wc.client, err = waddell.ConnectTLS(wc.conn, wc.serverCert)
	} else {
		wc.client, err = waddell.Connect(wc.conn)
	}
	if err == nil {
		log.Debugf("Connected to Waddell!! Id is: %s", wc.client.ID())
	} else {
		log.Debugf("Unable to connect waddell client: %s", err)
	}
	return
}

func (wc *waddellConn) close() error {
	return wc.conn.Close()
}

func idToBytes(id uint32) []byte {
	b := make([]byte, 4)
	endianness.PutUint32(b[:4], id)
	return b
}

// DefaultWaddellCert is the certificate for the production waddell server(s)
// used by, amongst other things, flashlight.
const DefaultWaddellCert = `-----BEGIN CERTIFICATE-----
MIIDkTCCAnmgAwIBAgIJAJKSxfu1psP7MA0GCSqGSIb3DQEBBQUAMF8xCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApTb21lLVN0YXRlMSkwJwYDVQQKDCBCcmF2ZSBOZXcg
U29mdHdhcmUgUHJvamVjdCwgSW5jLjEQMA4GA1UEAwwHd2FkZGVsbDAeFw0xNDEx
MDcyMDI5MDRaFw0xNTExMDcyMDI5MDRaMF8xCzAJBgNVBAYTAlVTMRMwEQYDVQQI
DApTb21lLVN0YXRlMSkwJwYDVQQKDCBCcmF2ZSBOZXcgU29mdHdhcmUgUHJvamVj
dCwgSW5jLjEQMA4GA1UEAwwHd2FkZGVsbDCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBAOz22kAZXaVmFzo8+qaYbDyiZSc+D6j4+uQDlCFYsymdMSBaMRho
D3HNXAuvlmYGvZIc/jCM0LJ8m0MjS8DDa/EOWBDNcLV9ABxfqxPaAm2u8EU8vP8G
E3eGmoSrD0tB/OAF/utFvAEPNShwhMc2aY4qWPPrNqWa5U8f0JLnoZbnOWxMteU7
uSC+pRUbl3+tueWvFr+hXZMuGzb2Mes0UapJ//BKbaz0XboQ9Y7cRj8OiXjh3x4K
4Rz9qN8CrgOtwL9HNJ6krcgwaYIrf8O14Acc8VzcASLdtwEerHWgm2EZG+FZ24yP
ZwDLlcxJul29gjGnVpxDJaeB/1P18680fKECAwEAAaNQME4wHQYDVR0OBBYEFC9r
MKrgfqko3g/n8fgg3PUq7UCTMB8GA1UdIwQYMBaAFC9rMKrgfqko3g/n8fgg3PUq
7UCTMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggEBAGlC2BrXcLZefm7G
IAZUjSj3nEPmoARH9Y2lxR78/FtAXu3WwXFeDY5wq1HRDWMUB/usBNk+19SXQjxF
ykZGqc5on7QSqbu489Kh37Jenfi6MGXPFh1brFaNuCndW3x/x2wer+k/y7HAXTN0
OGRaZaCqwkFoI0GCnmJUJdA1ahBYMkqFcWrvuw4aDzvfWYCfFUmADQdkb+xJGGiF
plISFgS/kDrK3OfIBu8S+XuhAIlzXKnHb+887pcvpm4f3zVv7rB8amv10x2E5fjS
RnCIHFZ4k1Au8N60Da3Z28hizafJeV4uHbzjYU+n8XpVqFqJI83CdsbiZ+nXO87G
pUWu27U=
-----END CERTIFICATE-----`

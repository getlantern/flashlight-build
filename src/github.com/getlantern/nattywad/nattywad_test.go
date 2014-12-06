package nattywad

import (
	"net"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/getlantern/waddell"
)

const (
	waddellAddr = "128.199.130.61:443"
)

// TestRoundTrip is an integration test that tests a round trip with client and
// server, using a waddell server in the cloud.
func TestRoundTrip(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	serverIdCh := make(chan waddell.PeerId)

	wc, err := waddell.NewClient(&waddell.ClientConfig{
		Dial: func() (net.Conn, error) {
			return net.Dial("tcp", waddellAddr)
		},
		ServerCert: DefaultWaddellCert,
		OnId: func(id waddell.PeerId) {
			serverIdCh <- id
		},
	})
	if err != nil {
		t.Fatalf("Unable to start waddell client: %s", err)
	}

	server := &Server{
		Client: wc,
		OnSuccess: func(local *net.UDPAddr, remote *net.UDPAddr) bool {
			log.Debugf("Success! %s -> %s", local, remote)
			wg.Done()
			return true
		},
		OnFailure: func(err error) {
			t.Errorf("Server - Traversal failed: %s", err)
			wg.Done()
		},
	}
	server.Start()

	clientMgr := &waddell.ClientMgr{
		Dial: func(addr string) (net.Conn, error) {
			return net.Dial("tcp", addr)
		},
		ServerCert:        DefaultWaddellCert,
		ReconnectAttempts: 10,
	}

	client := &Client{
		ClientMgr: clientMgr,
		OnSuccess: func(info *TraversalInfo) {
			log.Debugf("Client - Success! %s", spew.Sdump(info))
			wg.Done()
		},
		OnFailure: func(info *TraversalInfo) {
			t.Errorf("Client - Traversal failed: %s", spew.Sdump(info))
			wg.Done()
		},
	}
	serverId := <-serverIdCh
	client.Configure([]*ServerPeer{&ServerPeer{
		ID:          serverId.String(),
		WaddellAddr: waddellAddr,
	}})

	wg.Wait()
}

// This file is used to provide nattywad functionality on most systems.
//
// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris windows

package server

import (
	"github.com/getlantern/flashlight/globals"
	"github.com/getlantern/flashlight/nattest"
	"github.com/getlantern/nattywad"
	"github.com/getlantern/waddell"
	"net"
)

// startNattywad stars a nattywad server to provide NAT traversal
// functionality.
func (server *Server) startNattywad(waddellAddr string) {
	log.Debugf("Connecting to waddell at: %s", waddellAddr)
	var err error
	server.waddellClient, err = waddell.NewClient(&waddell.ClientConfig{
		Dial: func() (net.Conn, error) {
			return net.Dial("tcp", waddellAddr)
		},
		ServerCert:        globals.WaddellCert,
		ReconnectAttempts: 10,
		OnId: func(id waddell.PeerId) {
			log.Debugf("Connected to Waddell!! Id is: %s", id)
		},
	})
	if err != nil {
		log.Errorf("Unable to connect to waddell: %s", err)
		server.waddellClient = nil
		return
	}
	server.nattywadServer = &nattywad.Server{
		Client: server.waddellClient,
		OnSuccess: func(local *net.UDPAddr, remote *net.UDPAddr) bool {
			err := nattest.Serve(local)
			if err != nil {
				log.Error(err.Error())
				return false
			}
			return true
		},
	}
	server.nattywadServer.Start()
}

// stopNattywad stops the nattywad server.
func (server *Server) stopNattywad() {
	log.Debug("Stopping nattywad server")
	server.nattywadServer.Stop()
	server.nattywadServer = nil
	log.Debug("Stopping waddell client")
	server.waddellClient.Close()
	server.waddellClient = nil
}

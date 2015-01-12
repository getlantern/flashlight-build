package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/getlantern/fronted"
	"github.com/getlantern/golog"
	"github.com/getlantern/nattywad"
	"github.com/getlantern/waddell"

	"github.com/getlantern/flashlight/globals"
	"github.com/getlantern/flashlight/statreporter"
	"github.com/getlantern/flashlight/statserver"
)

const (
	PortmapFailure = 50
)

var (
	log = golog.LoggerFor("flashlight.server")
)

type Server struct {
	// Addr: listen address in form of host:port
	Addr string

	// Host: FQDN that is guaranteed to hit this server
	Host string

	// ReadTimeout: (optional) timeout for read ops
	ReadTimeout time.Duration

	// WriteTimeout: (optional) timeout for write ops
	WriteTimeout time.Duration

	CertContext                *fronted.CertContext // context for certificate management
	AllowNonGlobalDestinations bool                 // if true, requests to LAN, Loopback, etc. will be allowed

	waddellClient  *waddell.Client
	nattywadServer *nattywad.Server
	cfg            *ServerConfig
	cfgMutex       sync.Mutex
}

func (server *Server) Configure(newCfg *ServerConfig) {
	server.cfgMutex.Lock()
	defer server.cfgMutex.Unlock()

	oldCfg := server.cfg

	log.Debug("Server.Configure() called")
	if oldCfg != nil && reflect.DeepEqual(oldCfg, newCfg) {
		log.Debugf("Server configuration unchanged")
		return
	}

	if oldCfg == nil || newCfg.Portmap != oldCfg.Portmap {
		// Portmap changed
		if oldCfg != nil && oldCfg.Portmap > 0 {
			log.Debugf("Attempting to unmap old external port %d", oldCfg.Portmap)
			err := unmapPort(oldCfg.Portmap)
			if err != nil {
				log.Errorf("Unable to unmap old external port: %s", err)
			}
			log.Debugf("Unmapped old external port %d", oldCfg.Portmap)
		}

		if newCfg.Portmap > 0 {
			log.Debugf("Attempting to map new external port %d", newCfg.Portmap)
			err := mapPort(server.Addr, newCfg.Portmap)
			if err != nil {
				log.Errorf("Unable to map new external port: %s", err)
				os.Exit(PortmapFailure)
			}
			log.Debugf("Mapped new external port %d", newCfg.Portmap)
		}
	}

	nattywadIsEnabled := newCfg.WaddellAddr != ""
	nattywadWasEnabled := server.nattywadServer != nil
	waddellAddrChanged := oldCfg == nil && newCfg.WaddellAddr != "" || oldCfg != nil && oldCfg.WaddellAddr != newCfg.WaddellAddr

	if waddellAddrChanged {
		if nattywadWasEnabled {
			server.stopNattywad()
		}
		if nattywadIsEnabled {
			server.startNattywad(newCfg.WaddellAddr)
		}
	}

	server.cfg = newCfg
}

func (server *Server) ListenAndServe() error {
	if server.Host != "" {
		log.Debugf("Running as host %s", server.Host)
	}

	fs := &fronted.Server{
		Addr:                       server.Addr,
		Host:                       server.Host,
		ReadTimeout:                server.ReadTimeout,
		WriteTimeout:               server.WriteTimeout,
		CertContext:                server.CertContext,
		AllowNonGlobalDestinations: server.AllowNonGlobalDestinations,
	}

	if server.cfg.Unencrypted {
		log.Debug("Running in unencrypted mode")
		fs.CertContext = nil
	}

	// Add callbacks to track bytes given
	fs.OnBytesReceived = func(ip string, destAddr string, req *http.Request, bytes int64) {
		onBytesGiven(destAddr, req, bytes)
		statserver.OnBytesReceived(ip, bytes)
	}
	fs.OnBytesSent = func(ip string, destAddr string, req *http.Request, bytes int64) {
		onBytesGiven(destAddr, req, bytes)
		statserver.OnBytesSent(ip, bytes)
	}

	l, err := fs.Listen()
	if err != nil {
		return fmt.Errorf("Unable to listen at %s: %s", server.Addr, err)
	}
	return fs.Serve(l)
}

// determineInternalIP determines the internal IP to use for mapping ports. It
// does this by dialing a website on the public Internet and then finding out
// the LocalAddr for the corresponding connection. This gives us an interface
// that we know has Internet access, which makes it suitable for port mapping.
func determineInternalIP() (string, error) {
	conn, err := net.DialTimeout("tcp", "s3.amazonaws.com:443", 20*time.Second)
	if err != nil {
		return "", fmt.Errorf("Unable to determine local IP: %s", err)
	}
	defer conn.Close()
	host, _, err := net.SplitHostPort(conn.LocalAddr().String())
	return host, err
}

func onBytesGiven(destAddr string, req *http.Request, bytes int64) {
	_, port, _ := net.SplitHostPort(destAddr)
	if port == "" {
		port = "0"
	}

	given := statreporter.CountryDim().
		And("flserver", globals.InstanceId).
		And("destport", port)
	given.Increment("bytesGiven").Add(bytes)
	given.Increment("bytesGivenByFlashlight").Add(bytes)

	clientCountry := req.Header.Get("Cf-Ipcountry")
	if clientCountry != "" {
		givenTo := statreporter.Country(clientCountry)
		givenTo.Increment("bytesGivenTo").Add(bytes)
		givenTo.Increment("bytesGivenToByFlashlight").Add(bytes)
	}
}

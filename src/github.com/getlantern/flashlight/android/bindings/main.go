// package flashlight provides minimal configuration for spawning a flashlight
// client.

package flashlight

import (
	"github.com/getlantern/flashlight/client"
	"github.com/getlantern/flashlight/config"
	"github.com/getlantern/golog"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	log = golog.LoggerFor("flashlight")
)

// StopClientProxy should stop the client.
func StopClientProxy() {
	// TODO: We only want to stop the proxy, not to bring down the whole process.
	os.Exit(0)
}

// RunClientProxy creates a new client at the given address. If an active
// client is found it kill the client before starting a new one.
func RunClientProxy(listenAddr string) (err error) {
	var cfg *config.Config

	cfg = new(config.Config)
	cfg = &config.Config{
		Role:   "client",
		Client: &client.ClientConfig{},
		Addr:   listenAddr,
	}

	cfg.ApplyDefaults()

	client := &client.Client{
		Addr:         cfg.Addr,
		ReadTimeout:  0, // don't timeout
		WriteTimeout: 0,
	}

	client.Configure(cfg.Client)

	if err = client.ListenAndServe(); err != nil {
		panic(err.Error())
	}

	return nil
}

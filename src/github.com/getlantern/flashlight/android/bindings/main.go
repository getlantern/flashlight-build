// package bindings provides minimal configuration for spawning a flashlight
// client.

package bindings

import (
	"github.com/getlantern/flashlight/client"
	"github.com/getlantern/flashlight/config"
	"github.com/getlantern/flashlight/statreporter"
	"github.com/getlantern/flashlight/statserver"
	"github.com/getlantern/golog"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	log           = golog.LoggerFor("flashlight")
	cfg           *config.Config
	configUpdates = make(chan *config.Config)
)

func init() {
	var err error
	cfg, err = config.Start(func(updated *config.Config) {
		configUpdates <- updated
	})
	if err != nil {
		log.Fatalf("Unable to start configuration: %s", err)
	}

	configureStats(cfg, true)
}

func RunClientProxy(listenAddr string) (err error) {

	cfg.Addr = listenAddr

	client := &client.Client{
		Addr:         cfg.Addr,
		ReadTimeout:  0, // don't timeout
		WriteTimeout: 0,
	}

	client.Configure(cfg.Client)

	// Continually poll for config updates and update client accordingly
	go func() {
		for {
			cfg := <-configUpdates
			// TODO: We don't need this yet, this is a PoC.
			configureStats(cfg, false)

			client.Configure(cfg.Client)
		}
	}()

	if err = client.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func configureStats(cfg *config.Config, failOnError bool) {
	err := statreporter.Configure(cfg.Stats)
	if err != nil {
		log.Error(err)
		if failOnError {
			log.Fatalf("Config error.")
		}
	}

	if cfg.StatsAddr != "" {
		statserver.Start(cfg.StatsAddr)
	} else {
		statserver.Stop()
	}
}

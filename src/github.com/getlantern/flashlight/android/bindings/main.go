// package bindings provides minimal configuration for spawning a flashlight
// client.

package bindings

import (
	// "github.com/getlantern/flashlight/statreporter"
	// "github.com/getlantern/flashlight/statserver"
	"github.com/getlantern/flashlight/client"
	"github.com/getlantern/flashlight/config"
	"github.com/getlantern/golog"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	log = golog.LoggerFor("flashlight")
)

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
		return err
	}

	return nil
}

// func configureStats(cfg *config.Config, failOnError bool) {
// 	err := statreporter.Configure(cfg.Stats)
// 	if err != nil {
// 		log.Error(err)
// 		if failOnError {
// 			log.Fatalf("Config error.")
// 		}
// 	}
//
// 	if cfg.StatsAddr != "" {
// 		statserver.Start(cfg.StatsAddr)
// 	} else {
// 		statserver.Stop()
// 	}
// }

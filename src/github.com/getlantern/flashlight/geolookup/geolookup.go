package geolookup

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/getlantern/flashlight/ui"
	"github.com/getlantern/geolookup"
	"github.com/getlantern/golog"
)

const (
	sleepTime = time.Second * 10

	endpoint = "/geolookup"
)

var (
	log = golog.LoggerFor("geolookup-flashlight")

	startMutex sync.Mutex
	uichannel  *ui.UIChannel
	clientCity = make(chan *geolookup.City)
)

func watchForUserIP() {
	for {
		if !geolookup.UsesDefaultHTTPClient() {
			// Will look up only if we're using a proxy.
			geodata, err := geolookup.LookupCity("")
			if err == nil {
				clientCity <- geodata
				// We got what we wanted, no need to query for it again, let's exit.
				return
			}
		}
		// Sleep if the proxy is not ready yet of any error happened.
		time.Sleep(sleepTime)
	}
}

func StartService() error {
	// Looking up client's information.
	go watchForUserIP()

	startMutex.Lock()

	if uichannel == nil {
		start()
	} else {
		var b []byte
		var err error

		// Waiting for the city to be discovered.
		city := <-clientCity

		if b, err = json.Marshal(city); err != nil {
			return fmt.Errorf("Unable to marshal geolocation information: %q", err)
		}

		// Writing data to channel.
		uichannel.Out <- b
	}

	startMutex.Unlock()

	return nil
}

func start() {
	uichannel = ui.NewChannel(endpoint, func(func([]byte) error) error {
		// We don't really need to to anything here yet but we could add code for
		// allowing looking up other IPs.
		return nil
	})

	log.Debugf("Serving geodata information at %v", uichannel.URL)
}

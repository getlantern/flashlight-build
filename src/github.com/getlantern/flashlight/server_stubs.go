// +build android

package main

import (
	"runtime"

	"github.com/getlantern/flashlight/config"
)

// runServerProxy is not implemented.
func runServerProxy(cfg *config.Config) {
	log.Debugf("runServerProxy not implemented in %s/%s.", runtime.GOOS, runtime.GOARCH)
}

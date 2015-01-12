// This file is used to disable runServerProxy on android hosts.
//
// +build android

package main

import (
	"runtime"

	"github.com/getlantern/flashlight/config"
)

// runServerProxy is not implemented on android.
func runServerProxy(cfg *config.Config) {
	log.Debugf("runServerProxy not implemented in %s/%s.", runtime.GOOS, runtime.GOARCH)
}

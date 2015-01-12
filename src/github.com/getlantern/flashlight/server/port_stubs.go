// This file is used to disable igdman on android hosts.
//
// +build android

package server

import (
	"errors"
)

var (
	errNotImplemented = errors.New("Not implemented.")
)

// mapPort is not implemented on android.
func mapPort(addr string, port int) error {
	return errNotImplemented
}

// unmapPort is not implemented on android.
func unmapPort(port int) error {
	return errNotImplemented
}

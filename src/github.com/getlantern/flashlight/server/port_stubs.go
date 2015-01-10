// Not supported in android.
// +build android

package server

import (
	"errors"
)

var (
	errNotImplemented = errors.New("Not implemented.")
)

func mapPort(addr string, port int) error {
	return errNotImplemented
}

func unmapPort(port int) error {
	return errNotImplemented
}

// This file is used to disable natty functionality on android hosts.
//
// +build android

package bin

import (
	"errors"
)

var (
	errNotImplemented = errors.New(`Not implemented.`)
)

// Asset is not implemented on android.
func Asset(name string) ([]byte, error) {
	return nil, errNotImplemented
}

// AssetNames is not implemented on android.
func AssetNames() []string {
	return nil
}

// AssetDir is not implemented on android.
func AssetDir(name string) ([]string, error) {
	return nil, errNotImplemented
}

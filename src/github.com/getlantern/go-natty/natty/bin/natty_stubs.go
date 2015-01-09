// +build android

package bin

import (
	"errors"
)

var (
	errNotImplemented = errors.New(`Not implemented.`)
)

// Asset is not implemented.
func Asset(name string) ([]byte, error) {
	return nil, errNotImplemented
}

// AssetNames is not implemented.
func AssetNames() []string {
	return nil
}

func AssetDir(name string) ([]string, error) {
	return nil, errNotImplemented
}

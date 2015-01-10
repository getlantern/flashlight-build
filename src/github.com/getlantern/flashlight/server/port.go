// +build windows darwin freebsd openbsd !android,linux

package server

import (
	"fmt"
	"github.com/getlantern/go-igdman/igdman"
	"net"
	"strconv"
)

func mapPort(addr string, port int) error {
	internalIP, internalPortString, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("Unable to split host and port for %v: %v", addr, err)
	}

	internalPort, err := strconv.Atoi(internalPortString)
	if err != nil {
		return fmt.Errorf("Unable to parse local port: ")
	}

	if internalIP == "" {
		internalIP, err = determineInternalIP()
		if err != nil {
			return fmt.Errorf("Unable to determine internal IP: %s", err)
		}
	}

	igd, err := igdman.NewIGD()
	if err != nil {
		return fmt.Errorf("Unable to get IGD: %s", err)
	}

	igd.RemovePortMapping(igdman.TCP, port)
	err = igd.AddPortMapping(igdman.TCP, internalIP, internalPort, port, 0)
	if err != nil {
		return fmt.Errorf("Unable to map port with igdman %d: %s", port, err)
	}

	return nil
}

func unmapPort(port int) error {
	igd, err := igdman.NewIGD()
	if err != nil {
		return fmt.Errorf("Unable to get IGD: %s", err)
	}

	igd.RemovePortMapping(igdman.TCP, port)
	if err != nil {
		return fmt.Errorf("Unable to unmap port with igdman %d: %s", port, err)
	}

	return nil
}

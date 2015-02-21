// +build !windows

package runhide

import (
	"os/exec"
)

// Just use os/exec
func command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

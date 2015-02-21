// runhide provides a way to execute a program on Windows without it spawning
// a visible shell window. This is important when compiling a program with the
// linkger flag -H=windowsgui to build a gui-mode windows executable.
//
// On platforms other than Windows, it just passes through to exec.Command().
//
// It uses the Visual Basic code from
// http://www.robvanderwoude.com/files/runnhide_vbs.txt to actually do the
// running.
//
// See also http://stackoverflow.com/questions/3677773/how-can-i-run-a-windows-batch-file-but-hide-the-command-window.
package runhide

import (
	"os/exec"
)

func Command(name string, arg ...string) *exec.Cmd {
	return command(name, arg...)
}

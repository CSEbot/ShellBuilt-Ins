// bash.go

package builtins

import (
	"io"
	"os"
	"os/exec"
)

func RunBash(w io.Writer) error {
	// Start a new bash shell.
	cmd := exec.Command("bash")

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// Execute the command and return the error.
	return cmd.Run()
}

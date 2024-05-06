// sh.go

package builtins

import (
	"io"
	"os"
	"os/exec"
)

func RunShell(w io.Writer) error {
	// Start a new shell.
	cmd := exec.Command("sh")

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// Execute the command and return the error.
	return cmd.Run()
}

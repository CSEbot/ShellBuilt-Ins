// sh_test.go

package builtins_test

import (
    "bytes"
    "testing"

    "github.com/CSEbot/ShellBuilt-Ins/Project2/builtins"
    "github.com/stretchr/testify/require"
)

func TestRunShell(t *testing.T) {
    t.Parallel()

    w := &bytes.Buffer{}

    // Execute the shell command
    err := builtins.RunShell(w)
    require.NoError(t, err)

    // Since sh opens a new shell, it's difficult to test its output directly.
    // We can only test if it executes without errors.
}

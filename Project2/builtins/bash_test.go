// bash_test.go

package builtins_test

import (
    "bytes"
    "testing"

    "github.com/CSEbot/ShellBuilt-Ins/Project2/builtins"
    "github.com/stretchr/testify/require"
)

func TestRunBash(t *testing.T) {
    t.Parallel()

    w := &bytes.Buffer{}

    // Execute the bash command
    err := builtins.RunBash(w)
    require.NoError(t, err)

    // Since bash opens a new shell, it's difficult to test its output directly.
    // We can only test if it executes without errors.
}

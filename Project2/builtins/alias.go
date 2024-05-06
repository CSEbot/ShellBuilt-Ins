// alias.go

package builtins

import (
	"fmt"
	"io"
	"strings"
)

var aliases map[string]string

func init() {
	aliases = make(map[string]string)
}

// AliasCommand sets or displays aliases.
func AliasCommand(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// Display all aliases if no arguments provided
		for alias, command := range aliases {
			fmt.Fprintf(w, "%s='%s'\n", alias, command)
		}
		return nil
	}

	// Parse arguments as alias=value pairs
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			return fmt.Errorf("invalid alias format: %s", arg)
		}
		alias, command := parts[0], parts[1]
		aliases[alias] = command
	}

	return nil
}

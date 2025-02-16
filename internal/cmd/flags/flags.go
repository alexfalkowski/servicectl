package flags

import (
	"github.com/alexfalkowski/go-service/cmd"
)

// IsSet checks if a flag is set by name.
func IsSet(set *cmd.FlagSet, name string) bool {
	b, _ := set.GetBool(name)

	return b
}

package flags

import (
	"github.com/alexfalkowski/go-service/flags"
)

// IsSet checks if a flag is set by name.
func IsSet(set *flags.FlagSet, name string) bool {
	b, _ := set.GetBool(name)

	return b
}

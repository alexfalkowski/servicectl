package os

import (
	"os"

	"github.com/alexfalkowski/go-service/runtime"
)

// WriteFile to path with data.
func WriteFile(path string, data []byte) {
	path += "-new"

	err := os.WriteFile(path, data, 0o600)
	runtime.Must(err)
}

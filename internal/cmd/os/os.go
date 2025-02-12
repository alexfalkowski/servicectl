package os

import (
	"encoding/base64"
	"os"
)

// WriteFile to path with data.
func WriteFile(path string, data []byte) error {
	path += "-new"

	return os.WriteFile(path, data, 0o600)
}

// WriteBase64File to path with data.
func WriteBase64File(path string, data []byte) error {
	return WriteFile(path, []byte(base64.StdEncoding.EncodeToString(data)))
}

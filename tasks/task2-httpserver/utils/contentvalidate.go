// contentvalidate.go
package utils

import (
	"bytes"
	"errors"
)

// IsEmptyContent reports true if b is empty or only whitespace.
func IsEmptyContent(b []byte) bool {
	return len(bytes.TrimSpace(b)) == 0
}

// RequireNonEmpty returns an error if b is empty or only whitespace.
func RequireNonEmpty(b []byte) error {
	if IsEmptyContent(b) {
		return errors.New("empty content")
	}
	return nil
}

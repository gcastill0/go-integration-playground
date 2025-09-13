// jsonvalidate.go
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ValidateJSONBytes returns nil iff data is well-formed JSON and contains
// exactly one top-level value (object, array, string, number, true/false/null).
func ValidateJSONBytes(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber() // avoid float rounding for large ints during parsing

	var v any
	if err := dec.Decode(&v); err != nil {
		// Enhance syntax errors with position (byte offset).
		var syn *json.SyntaxError
		if errors.As(err, &syn) {
			return fmt.Errorf("syntax error at byte %d: %v", syn.Offset, err)
		}
		return err
	}

	// Ensure no non-whitespace remains (i.e., no second top-level value).
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return errors.New("extra data after top-level value")
		}
		return err
	}
	return nil
}

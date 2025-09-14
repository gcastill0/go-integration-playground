package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Preflight checks that url is reachable and readable.
// It succeeds on 2xx/3xx; tries HEAD first, then falls back to GET Range 0-0.
func Preflight(url string) error {

	hc := &http.Client{Timeout: 10 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try HEAD
	if req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil); err == nil {
		if resp, err := hc.Do(req); err == nil {
			defer resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 400 {
				return nil
			}
			// Only fall back if HEAD isn’t supported.
			if resp.StatusCode != http.StatusMethodNotAllowed && resp.StatusCode != http.StatusNotImplemented {
				return fmt.Errorf("preflight HEAD: %s", resp.Status)
			}
		}
	}

	// Fallback: GET a single byte
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", "bytes=0-0")

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("preflight GET: %s", resp.Status)
	}
	_, _ = io.CopyN(io.Discard, resp.Body, 1)
	return nil
}

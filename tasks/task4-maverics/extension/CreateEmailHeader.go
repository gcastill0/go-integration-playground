package CreateEmailHeader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/strata-io/service-extension/orchestrator"
)

// CreateEmailHeader creates an email header from the user information extracted in Task 3.
func CreateEmailHeader(
	api orchestrator.Orchestrator,
	_ http.ResponseWriter,
	_ *http.Request,
) (http.Header, error) {
	logger := api.Logger()
	logger.Info("se", "building email custom header")

	// Ensure we have an orchestrator session.
	session, err := api.Session()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Orchestrator session: %w", err)
	}
	_ = session // (not used further; keep to satisfy the compiler)

	logger.Debug("se", "retrieving email from mock endpoint..")

	// Configure the mock endpoint: allow override via env, otherwise use JSONPlaceholder user #2.
	url := os.Getenv("MOCK_USER_URL")
	if url == "" {
		url = "https://jsonplaceholder.typicode.com/users/2"
	}

	// Fetch the user’s email address from the mock endpoint. Check for errors.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch email: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch email: unexpected status: %s", resp.Status)
	}

	// Unmarshal the JSON response.
	var body struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	// Set the user’s email address.
	email := strings.TrimSpace(body.Email)
	if email == "" {
		return nil, fmt.Errorf("email missing in response")
	}

	// Build and return the header to the orchestrator.
	header := make(http.Header)
	header["CUSTOM-EMAIL"] = []string{email}
	return header, nil
}

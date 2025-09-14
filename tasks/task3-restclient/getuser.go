// /tasks/task3-restclient/getuser.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Geo struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

type User struct {
	ID int `json:"id"`
	// Name     string  `json:"name"`
	// UserName string  `json:"username"`
	Email string `json:"email"`
	// Address  Address `json:"address"`
	Phone string `json:"phone"`
	// Website  string  `json:"website"`
	Company Company `json:"company"`
}

// GetUser initializes an HTTP client, calls the REST endpoint,
// unmarshals the JSON response into User, prints it, and returns it.
func GetUser(baseURL string) (*User, error) {
	// initialize HTTP client
	hc := &http.Client{Timeout: 10 * time.Second}

	// Make a call to the REST endpoint
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	// Unmarshal the JSON response
	var u User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	// Print JSON contents (pretty)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(u); err != nil {
		return nil, fmt.Errorf("print: %w", err)
	}

	return &u, nil
}

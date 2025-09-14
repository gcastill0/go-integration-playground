// tasks/task2-httpserver/main.go
// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2025 G Castillo

// Command task2-httpserver starts a minimal HTTP server that listens on
// port 8080 and serves a plain-text greeting at the index path ("/").
// It is intentionally small for demonstration and bootstrap purposes.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	emailheader "github.com/gcastillo/go-integration-playground/tasks/task4-maverics/extension"
	"github.com/strata-io/service-extension/orchestrator"
)

// ---- ultra-minimal fake api (only what CreateEmailHeader uses) ----
type nopLogger struct{}

func (nopLogger) Debug(...any) {}
func (nopLogger) Info(...any)  {}
func (nopLogger) Error(...any) {}

type fakeAPI struct{}

func (fakeAPI) Logger(...any) any           { return nopLogger{} }
func (fakeAPI) Session(...any) (any, error) { return struct{}{}, nil }

var _ orchestrator.Orchestrator = (*fakeAPI)(nil) // keep your full fake if needed

// main starts a simple HTTP server on port 8080.
// It registers a handler for the root path ("/") that responds with "Hello worlk opportunity".
// The server logs any fatal errors encountered during startup or execution.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// minimal index response

		api := &fakeAPI{}

		hdr, err := emailheader.CreateEmailHeader(api, w, r)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// show what the upstream would see (for easy testing)
		type out struct {
			Headers http.Header `json:"headers"`
		}

		json.NewEncoder(w).Encode(out{Headers: hdr})

		fmt.Fprintln(w, "Hello worlk opportunity")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

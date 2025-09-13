// tasks/task2-httpserver/main.go
// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2024 G Castillo

// Command task2-httpserver starts a minimal HTTP server that listens on
// port 8080 and serves a plain-text greeting at the index path ("/").
// It is intentionally small for demonstration and bootstrap purposes.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	utils "github.com/gcastill0/go-integration-playground/tasks/task2-httpserver/utils"
)

// main starts a simple HTTP server on port 8080.
// It registers a handler for the root path ("/") that responds with "Hello, world".
// The server logs any fatal errors encountered during startup or execution.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// minimal index response
		fmt.Fprintln(w, "Hello, world")
	})

	// /ping: JSON pong
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"pong"}`))
	})

	// /echo: accept JSON and return it unchanged
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Reject empty/whitespace-only payloads up front
		if err := utils.RequireNonEmpty(body); err != nil {
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		// Validate Content-Type == application/json here.
		if err := utils.ValidateJSONBytes(body); err != nil {
			http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

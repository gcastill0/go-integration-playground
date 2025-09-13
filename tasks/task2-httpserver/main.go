// tasks/task2-httpserver/main.go
// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2024 G Castillo

// Command task2-httpserver starts a minimal HTTP server that listens on
// port 8080 and serves a plain-text greeting at the index path ("/").
// It is intentionally small for demonstration and bootstrap purposes.

package main

import (
	"fmt"
	"log"
	"net/http"
)

// main starts a simple HTTP server on port 8080.
// It registers a handler for the root path ("/") that responds with "Hello, world".
// The server logs any fatal errors encountered during startup or execution.

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// minimal index response
		fmt.Fprintln(w, "Hello, world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

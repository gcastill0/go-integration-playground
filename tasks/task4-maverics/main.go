// tasks/task4-maverics-ext/dev_harness/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/strata-io/service-extension/orchestrator"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		var api orchestrator.Orchestrator // nil is OK since SE doesn't use it here

		hdr, err := CreateEmailHeader(api, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// show the headers as JSON so it's easy to assert in tests
		_ = json.NewEncoder(w).Encode(map[string]any{"headers": hdr})
	})

	log.Println("dev harness running at http://127.0.0.1:8081/headers")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

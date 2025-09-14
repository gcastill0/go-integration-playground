// tasks/task3-restclient/main.go
package main

import (
	"fmt"
	"os"

	utils "github.com/gcastill0/go-integration-playground/tasks/task3-restclient/utils"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: %s <URL>\n", os.Args[0])
		os.Exit(2)
	}

	baseURL := os.Args[1]

	// Pre-flight
	if err := utils.Preflight(baseURL); err != nil {
		fmt.Fprintln(os.Stderr, "Preflight:", err)
		os.Exit(3)
	}

	GetUser(baseURL)
}

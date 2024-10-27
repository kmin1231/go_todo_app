// webserver.go
package main

import (
	"fmt"
	"net/http"  // HTTP clients & servers
	"os"
)

func main() {
	fmt.Printf("Attempting to connect the server...\n")

	err := http.ListenAndServe(
		":18080",  // 'localhost' by default (loopback); port number specified
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {  // handler
			// w: response writer
			// r: request object
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])  // sends a response to client
			// Fprintf: prints formatted strings to a specified io.Writer
			// ex. files, network connections, HTTP responses, etc.
		}),
	)  // starts server
	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)  // 'v' format (value)
		os.Exit(1)  // normal (0) vs. abnormal (1)
	}
}
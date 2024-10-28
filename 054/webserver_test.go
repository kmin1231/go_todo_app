// webserver_test.go (Code14.7)
package main

import "testing"

// defines testing function
func TestMainFunc(t *testing.T) {
	go main()   // goroutine
}
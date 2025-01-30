package main

// Import
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Configuration
var PORT int = 4000

// Main
func main() {
	// Def Routes
	http.HandleFunc("/hello", getHello)

	fmt.Printf("Starting server on port %d...\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

/**
  * Handler Func
*/
func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, Word HTTP!\n")
}
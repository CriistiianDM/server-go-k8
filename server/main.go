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
	http.HandleFunc("/healthz", healhzStatus)

	// Middleware CORS
	http.Handle("/", enableCORS(http.DefaultServeMux))

	fmt.Printf("Starting server on port %d...\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

/**
  * Cors Configuration
*/
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/**
  * Handler Func
*/
func getHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello Word!! IN GKE\n")
}

/**
  * Status of server
*/
func healhzStatus(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "Ok\n")
}
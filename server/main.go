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
	// Create a new ServeMux for routing
	mux := http.NewServeMux()

	// Def Routes
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/healthz", healhzStatus)
	mux.HandleFunc("/404", notFound)

	// Middleware CORS applied only to the mux
	http.Handle("/", enableCORS(mux))

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

/**
  * Handler for undefined routes
*/
func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "404 Not Found\n")
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "dev"

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	// Injected from the Kubernetes Secret that External Secrets Operator builds
	// from Secret Manager. Falls back so the binary still runs outside the
	// cluster, where nothing sets it.
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "(unset)"
	}
	fmt.Fprintf(w, "version=%s hostname=%s greeting=%s\n", version, hostname, greeting)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

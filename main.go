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
	fmt.Fprintf(w, "version=%s hostname=%s\n", version, hostname)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

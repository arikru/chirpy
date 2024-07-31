package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const filePath = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filePath))))
	mux.Handle("/app/assets/logo.png", http.FileServer(http.Dir(filePath+"/assets/logo.png")))

	healthzHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("OK"))
	}

	mux.HandleFunc("/healthz", healthzHandler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Serving files from %s on port: %v", filePath, port)
	log.Fatal(server.ListenAndServe())
}

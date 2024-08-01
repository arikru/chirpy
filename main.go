package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const filePath = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileServerHits: 0,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/*", http.StripPrefix("/app", apiCfg.middleWareMetricsInc(http.FileServer(http.Dir(filePath)))))
	// mux.Handle("/app/assets/logo.png", http.FileServer(http.Dir(filePath+"/assets/logo.png")))
	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("GET /metrics", apiCfg.handlerHits)
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Serving files  on http://localhost:%v/app", port)
	log.Fatal(server.ListenAndServe())

}

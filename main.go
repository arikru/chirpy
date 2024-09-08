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
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerHits)
	mux.HandleFunc("/api/reset", apiCfg.handlerReset)
	mux.HandleFunc("/api/validate_chirp", handlerValidate)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Serving files  on http://localhost:%v/app", port)
	log.Fatal(server.ListenAndServe())

}

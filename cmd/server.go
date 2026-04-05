package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := map[string]string{
			"status": "ok",
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to encode json")
		}
	})

	server := &http.Server{
		Handler: mux,
		Addr:    ":8000",
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "Server Failed: ", err)
	}
}

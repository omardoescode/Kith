package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"kith/internal/config"
)

func main() {
	var cfg config.Config
	err := cfg.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load config:", err)
		os.Exit(1)
	}

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
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	slog.Info(fmt.Sprintf("Running the project on port: %d", cfg.Port))
	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, "Server Failed: ", err)
	}
}

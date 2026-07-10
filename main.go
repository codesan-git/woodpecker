// Package main — Aplikasi demo sederhana untuk Woodpecker CI.
// Menyediakan HTTP server dengan endpoint /health dan frontend statis.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// HealthResponse adalah struktur response untuk endpoint /health
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

// healthHandler menangani request ke /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   getVersion(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func getVersion() string {
	if v := os.Getenv("APP_VERSION"); v != "" {
		return v
	}
	return "dev"
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", healthHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Printf("🚀 Server berjalan di http://0.0.0.0:%s", port)
	fmt.Printf("\n   🌐 Frontend : http://localhost:%s\n", port)
	fmt.Printf("   ❤️  Health   : http://localhost:%s/health\n\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

// New returns the root HTTP handler with all routes registered.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", health)
	mux.HandleFunc("GET /api/hello", hello)
	mux.HandleFunc("GET /api/message", getMessage)
	mux.HandleFunc("POST /api/message", postMessage)
	return cors(mux)
}

func health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func hello(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "Hello from the Go backend"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// cors allows the frontend origin (CORS_ORIGIN, default Vite dev server).
func cors(next http.Handler) http.Handler {
	origin := os.Getenv("CORS_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type messageBody struct {
	Message string `json:"message"`
}

// getMessage is a public endpoint that returns a fixed message.
func getMessage(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, messageBody{Message: "Hello! This is a message from the Go API."})
}

// postMessage is a public endpoint that accepts {"message": "..."} and sends it back.
func postMessage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var in messageBody
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Message == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "body must be JSON with a non-empty \"message\""})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "received", "message": in.Message})
}

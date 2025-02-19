package main

import (
    "log"
    "net/http"
    "gnark-auth/internal/auth"
)

func main() {
    server := auth.NewServer()

    http.HandleFunc("/authenticate", server.HandleAuthentication)
    
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
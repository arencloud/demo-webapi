package main

import (
    "fmt"
    "net/http"
)

func pingHandler(w http.ResponseWriter, _ *http.Request) {
    fmt.Fprintln(w, "Hello I am Ping")
}

func main() {
    http.HandleFunc("/api/ping", pingHandler)

    fmt.Println("Server starting on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server failed: %v\\n", err)
    }
}

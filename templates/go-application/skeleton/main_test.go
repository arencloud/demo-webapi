package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestPingHandler(t *testing.T) {
    req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "/api/ping", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(pingHandler)

    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
    }

    if !strings.Contains(rr.Body.String(), "Hello I am Ping") {
        t.Fatalf("handler returned unexpected body: %v", rr.Body.String())
    }
}

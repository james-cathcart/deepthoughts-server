package handler

import (
    "log"
    "net/http"
)

func DisableCors(h http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            log.Printf("disabling cors for %s request: %s\n", r.Method, r.URL)
            w.Header().Set("Access-Control-Allow-Origin", "*")
            h.ServeHTTP(w, r)
        },
    )
}

func JsonResponse(w http.ResponseWriter, status int, jsonBytes []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _, _ = w.Write(jsonBytes)
}
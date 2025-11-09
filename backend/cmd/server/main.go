package main

import (
    "fmt"
    "net/http"

    "example.com/backend/internal/math"
)

func handler(w http.ResponseWriter, r *http.Request) {
    sum := math.Add(2, 3)
    fmt.Fprintf(w, "sum=%d", sum)
}

func main() {
    http.HandleFunc("/health", handler)
    _ = http.ListenAndServe(":8080", nil)
}

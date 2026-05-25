package main

import (
	"jam/internal/rest"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/ping", rest.Ping)
	http.ListenAndServe(":8080", mux)
}

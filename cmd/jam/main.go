package main

import (
	"jam/internal/config"
	"jam/internal/rest"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/ping", rest.Ping)
	http.ListenAndServe(cfg.Host+":"+cfg.Port, mux)
}

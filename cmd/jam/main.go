package main

import (
	"jam/internal/config"
	"jam/internal/database"
	"jam/internal/rest"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Postgres\t%s@%s:%d\n", cfg.PostgresUser, cfg.PostgresHost, cfg.PostgresPort)
	db, err := database.Load(cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB, cfg.PostgresUser, cfg.PostgresPassword)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Printf("HTTP\t\t%s:%s\n", cfg.HTTPHost, cfg.HTTPPort)
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/ping", rest.Ping)
	http.ListenAndServe(cfg.HTTPHost+":"+cfg.HTTPPort, mux)
}

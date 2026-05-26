package database

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

func Load(host string, port uint16, database string, user string, password string) (*sql.DB, error) {
	config := pq.Config{
		Host:           host,
		Database:       database,
		Port:           port,
		User:           user,
		Password:       password,
		ConnectTimeout: 5 * time.Second,
		SSLMode:        "disable",
	}

	c, err := pq.NewConnectorConfig(config)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(c)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

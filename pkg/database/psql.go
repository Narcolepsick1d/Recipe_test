package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

var counts int64

func NewPostgresConnection(info ConnectionInfo) (*sql.DB, error) {
	for {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
			info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++

		} else {
			log.Println("Connected to Postgres!")
			return db, nil
		}
		if err := db.Ping(); err != nil {
			return nil, err
		}
		if counts > 10 {
			log.Println(err)
			return nil, err
		}
		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue

	}
}

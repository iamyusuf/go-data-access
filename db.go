package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type DBConfig struct {
	User   string
	Passwd string
}

func (c DBConfig) GetDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   c.User,
		Passwd: c.Passwd,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Capture connection properties.
	cfg := &DBConfig{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
	}

	db, err := cfg.GetDB()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}

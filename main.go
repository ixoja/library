package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config := Config{}
	config.WithFlags()

	db, err := sql.Open("sqlite3", "./shorten.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("failed to start db connection:", err.Error())
		}
	}()

	//st := storage.New(db)
	if err := st.InitDB(); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	s := service.New(controller.New(storage.NewCache(), st))
	lis, err := net.Listen("tcp", config.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

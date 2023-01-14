package main

import (
	"log"
	"settingsstore/src"
)

//go:generate rm -rf ../gen
//go:generate mkdir -p ../gen
//go:generate swagger generate server -t ../gen -f ../spec.yml --exclude-main --strict-responders
//go:generate go mod tidy

func main() {
	db, err := src.DatabaseSetup()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := src.ApiSetup(db)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

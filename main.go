package main

import (
	"fmt"
	"log"
	"settingsstore/src"
)

func main() {
	db, err := src.DatabaseSetup()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := src.ApiSetup(db)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("http://localhost:8080/docs")
	fmt.Println("http://localhost:8081/?pgsql=db&username=postgres&db=postgres&ns=public&select=settings")
	defer server.Shutdown()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

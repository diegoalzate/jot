package main

import (
	"jot/internal/server"
	"log"
)

func main() {
	srv := server.NewServer()
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

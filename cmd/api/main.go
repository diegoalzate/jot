package main

import (
	"fmt"

	"github.com/diegoalzate/jot/cmd/api/server"
)

func main() {
	srv, err := server.New()

	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	fmt.Println("listening on " + srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

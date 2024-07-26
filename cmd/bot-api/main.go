package main

import (
	"fmt"

	"github.com/diegoalzate/jot/internal/api"
)

func main() {
	srv, err := api.NewServer()

	if err != nil {
		panic(fmt.Sprintf("cannot create server: %s", err))
	}

	fmt.Println("listening on " + srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

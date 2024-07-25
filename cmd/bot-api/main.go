package main

import (
	"fmt"

	"github.com/diegoalzate/jot/internal/server"
)

func main() {
	srv := server.NewApiServer()
	fmt.Println("listening on " + srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

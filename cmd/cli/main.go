package main

import (
	"github.com/diegoalzate/jot/cmd/cli/cmd"
	"github.com/diegoalzate/jot/internal/database"
)

func main() {
	db := database.New()
	cmd.Execute(db)
}

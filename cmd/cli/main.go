package main

import (
	"github.com/diegoalzate/jot/cmd/cli/cmd"
	"github.com/diegoalzate/jot/internal/database"
)

func main() {
	db := database.New()
	// since this is cli we need to close this somehow
	defer db.Close()

	cmd.Execute(db)
}

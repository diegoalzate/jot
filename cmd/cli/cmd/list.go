package cmd

import (
	"log"

	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/spf13/cobra"
)

func newListCmd(db database.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			q := query.New(db.Conn)

			tasks, err := q.ListTasks(cmd.Context())

			if err != nil {
				log.Printf("failed to list tasks: %v", err)
			}

			for _, task := range tasks {
				cmd.Printf("- title: %s\n", task.Name)
			}
		},
	}

	return cmd
}

package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/diegoalzate/jot/cmd/cli/ui"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/diegoalzate/jot/internal/query"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newCreateCmd(db database.Service) *cobra.Command {
	var title string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "saves a task",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" {
				t := tea.NewProgram(ui.New(db))
				if _, err := t.Run(); err != nil {
					fmt.Printf("Alas, there's been an error: %v", err)
					os.Exit(1)
				}
			}

			q := query.New(db.Conn)

			_, err := q.CreateTask(cmd.Context(), query.CreateTaskParams{
				ID:          uuid.New().String(),
				Name:        title,
				Description: sql.NullString{},
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})

			if err != nil {
				log.Printf("failed to save task: %v", err)
			}
		},
	}

	// flags
	cmd.Flags().StringVarP(&title, "title", "t", "", "")

	return cmd
}

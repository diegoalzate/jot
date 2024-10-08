package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/diegoalzate/jot/cmd/cli/ui"
	"github.com/diegoalzate/jot/internal/database"
	"github.com/spf13/cobra"
)

var title string

func newCreateCmd(db database.Service) *cobra.Command {
	var title string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "saves a task",
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" {
				t := tea.NewProgram(ui.New())
				if _, err := t.Run(); err != nil {
					fmt.Printf("Alas, there's been an error: %v", err)
					os.Exit(1)
				}
			}

			// assume we have title and we will save this
			log.Print(title)
			return
		},
	}

	// flags
	cmd.Flags().StringVarP(&title, "title", "t", "", "")

	return cmd
}

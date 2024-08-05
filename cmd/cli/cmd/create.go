package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/diegoalzate/jot/cmd/cli/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "saves a task",
	Run: func(cmd *cobra.Command, args []string) {
		str := strings.Join(args, " ")

		if str == "" {
			t := tea.NewProgram(ui.New())
			if _, err := t.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
		}

		log.Print(str)
		return
	},
}

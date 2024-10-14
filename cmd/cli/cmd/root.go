package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/diegoalzate/jot/internal/database"
	"github.com/spf13/cobra"
)

func newRootCmd(db database.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "jot",
		Short: "Jot helps keep all your project thoughts in one place",
		Long:  "Jot should be able to connect to your project manager through its api and allow you to quickly access this in one line",

		Run: func(cmd *cobra.Command, args []string) {
			log.Print("cobra running")
		},
	}
}

func Execute(db database.Service) {
	rootCmd := newRootCmd(db)
	rootCmd.AddCommand(newCreateCmd(db))
	rootCmd.AddCommand(newListCmd(db))
	rootCmd.AddCommand(newVersionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

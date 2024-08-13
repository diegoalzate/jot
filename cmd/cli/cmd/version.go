package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `All software has versions. This is Jot's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Jot v0.1-beta")
		},
	}
}

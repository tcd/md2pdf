package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information about md2pdf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("md2pdf: v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information about md2pdf",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(m2p.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

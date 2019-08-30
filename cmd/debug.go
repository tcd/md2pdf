package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:    "debug [FILE]",
	Short:  "Output HTML & JSON information in addition to a PDF.",
	Hidden: true,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.Llongfile)
		err := m2p.Debug(args[0], "")
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}

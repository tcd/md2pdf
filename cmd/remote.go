package cmd

import (
	"github.com/spf13/cobra"
	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:    "remote [URL] [OUTPUT_NAME]",
	Short:  "Convert a remote markdown file to a PDF",
	Args:   cobra.ExactArgs(2),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		m2p.MdURLToPdfFile(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}

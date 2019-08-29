package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "md2pdf path/to/markdownfile.md",
	Short: "Generate PDFs from markdown files",
	Long:  `Generate PDFs from markdown files`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		if len(output) > 0 {
			rootFuncCustomOutput(args[0], output)
		}

		rootFuncDefault(args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "", "Write the resulting PDF to the named output file")
	// rootCmd.PersistentFlags().BoolP("silent", "s", false, "Only output error messages")
	// rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}

func rootFuncDefault(path string) {
	newFile, err := m2p.MdFileToUnnamedPdf(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PDF generated:", newFile)
	os.Exit(0)
}

func rootFuncCustomOutput(inPath, outPath string) {
	newFile, err := m2p.MdFileToPdfFile(inPath, outPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PDF generated:", newFile)
	os.Exit(0)
}

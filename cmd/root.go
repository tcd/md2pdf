package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "md2pdf path/to/markdownfile.md",
	Short: "Generate PDFs from markdown files.",
	Long:  `Generate PDFs from markdown files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		html, _ := cmd.Flags().GetBool("html")
		if html {
			inFile := args[0]
			inFile, outFile := parsePaths(inFile, ".html")
			err := m2p.Md2HTMLFile(inFile, outFile)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("HTML generated: ", outFile)
			os.Exit(0)
		} else {
			inFile := args[0]
			inFile, outFile := parsePaths(inFile, ".pdf")
			err := m2p.Md2PDF(inFile, outFile)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("HTML generated: ", outFile)
			os.Exit(0)
		}
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
	rootCmd.PersistentFlags().Bool("html", false, "Generate HTML output instead of PDF")
}

// parsePaths returns an absolute path to the given existing file,
// and the full path of the new file to create.
func parsePaths(inFile, ext string) (string, string) {
	oldFile := inFile
	if !filepath.IsAbs(oldFile) {
		absPath, err := filepath.Abs(oldFile)
		if err != nil {
			log.Fatal(err)
		}
		oldFile = absPath
	}

	baseName := filepath.Base(oldFile)
	extension := filepath.Ext(baseName)
	newFileName := strings.Replace(baseName, extension, ext, -1)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	newFile := filepath.Join(cwd, newFileName)
	return oldFile, newFile
}

package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	m2p "github.com/tcd/md2pdf/internal/md2pdf"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "md2pdf [FILE]",
	Short: "Generate PDFs from markdown files",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Run silently for --silent
		silent, err := cmd.Flags().GetBool("silent")
		logFatal(err)
		if silent {
			log.SetOutput(ioutil.Discard)
		}

		// Print Version for --version
		version, err := cmd.Flags().GetBool("version")
		logFatal(err)
		if version {
			log.Println(m2p.Version)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		outPath, _ := cmd.Flags().GetString("output")

		// Read from stdin if --stdin
		stdin, err := cmd.Flags().GetBool("stdin")
		logFatal(err)
		if stdin {
			bytes, err := ioutil.ReadAll(os.Stdin)
			logFatal(err)
			if string(bytes) != "" {
				newFile, err := m2p.MdBytesToPdfFile(bytes, outPath)
				logFatal(err)
				log.Println("PDF generated:", newFile)
				os.Exit(0)
			}
		}

		// Handle single file
		if len(args) == 1 {
			newFile, err := m2p.MdFileToPdfFile(args[0], outPath)
			logFatal(err)
			log.Println("PDF generated:", newFile)
			os.Exit(0)
		}

		cmd.Help()
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "", "Write the resulting PDF to the named output file")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "Silence all output messages")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
	rootCmd.Flags().Bool("stdin", false, "Read markdown content from stdin instead of a file")
}

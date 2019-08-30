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
		silent, err := cmd.Flags().GetBool("silent")
		if err != nil {
			log.Fatal(err)
		}
		if silent {
			log.SetOutput(ioutil.Discard)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		if version {
			log.Println(m2p.Version)
			os.Exit(0)
		}

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		outPath, _ := cmd.Flags().GetString("output")
		newFile, err := m2p.MdFileToPdfFile(args[0], outPath)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("PDF generated:", newFile)
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
}

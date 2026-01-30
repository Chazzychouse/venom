package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	logger  *log.Logger
)

var rootCmd = &cobra.Command{
	Use:   "venom",
	Short: "Audio engineering toolkit",
	Long:  `Venom is a CLI toolkit for audio file management, analysis, and organization.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Setup logger based on verbosity
		logger = log.New(os.Stderr)
		if verbose {
			logger.SetLevel(log.DebugLevel)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags available to all subcommands
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Register subcommands
	rootCmd.AddCommand(metadataCmd)
	rootCmd.AddCommand(bpmCmd)
	rootCmd.AddCommand(analyzeCmd)
	rootCmd.AddCommand(organizeCmd)
}

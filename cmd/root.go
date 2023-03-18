package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xapic",
	Short: "xapi compiler - a tool to validate, bundle and generate SDK from specification written using OpenAPI with custom extensions",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

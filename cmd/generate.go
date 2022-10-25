package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate SDK",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

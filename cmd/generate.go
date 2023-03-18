package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xapi-tools/xapic/pkg/gogen"
)

var generateCmd = &cobra.Command{
	Use:   "generate [spec path]",
	Short: "Generate SDK",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return gogen.ParseYamlPath(args[0])
	},
}

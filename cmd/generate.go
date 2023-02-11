package cmd

import (
	"github.com/ashutshkumr/openapiart/pkg/gogen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [spec path]",
	Short: "Generate SDK",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return gogen.ParseYamlPath(args[0])
	},
}

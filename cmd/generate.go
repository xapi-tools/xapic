package cmd

import (
	"github.com/ashutshkumr/openapiart/pkg/gogen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate SDK",
	RunE: func(cmd *cobra.Command, args []string) error {
		gogen.ParseYamlPath("api/openapi.yaml")
		return nil
	},
}

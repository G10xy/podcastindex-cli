package cmd

import (
	"github.com/spf13/cobra"
)

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "PodcastIndex categories",
}

var categoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all categories",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := newClient().CategoriesList(cmd.Context())
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

func init() {
	categoriesCmd.AddCommand(categoriesListCmd)
	rootCmd.AddCommand(categoriesCmd)
}

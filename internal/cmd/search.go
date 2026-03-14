package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/G10xy/podcastindex-cli/internal/client"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the PodcastIndex",
}

var searchByTermCmd = &cobra.Command{
	Use:   "byterm",
	Short: "Search podcasts by term",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}

		c := newClient()
		result, err := c.SearchByTerm(cmd.Context(), client.SearchByTermOptions{
			Query:    query,
			Val:      mustString(cmd, "val"),
			Max:      mustInt(cmd, "max"),
			APOnly:   mustBool(cmd, "aponly"),
			Clean:    mustBool(cmd, "clean"),
			Similar:  mustBool(cmd, "similar"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var searchByTitleCmd = &cobra.Command{
	Use:   "bytitle",
	Short: "Search podcasts by title",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}

		c := newClient()
		result, err := c.SearchByTitle(cmd.Context(), client.SearchByTitleOptions{
			Query:    query,
			Val:      mustString(cmd, "val"),
			Max:      mustInt(cmd, "max"),
			Clean:    mustBool(cmd, "clean"),
			Similar:  mustBool(cmd, "similar"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

var searchByPersonCmd = &cobra.Command{
	Use:   "byperson",
	Short: "Search episodes by person",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}

		c := newClient()
		result, err := c.SearchByPerson(cmd.Context(), client.SearchByPersonOptions{
			Query:    query,
			Max:      mustInt(cmd, "max"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Items))
	},
}

var searchMusicCmd = &cobra.Command{
	Use:   "music",
	Short: "Search music podcasts by term",
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}

		c := newClient()
		result, err := c.SearchMusicByTerm(cmd.Context(), client.SearchMusicByTermOptions{
			Query:    query,
			Val:      mustString(cmd, "val"),
			Max:      mustInt(cmd, "max"),
			APOnly:   mustBool(cmd, "aponly"),
			Clean:    mustBool(cmd, "clean"),
			FullText: mustBool(cmd, "fulltext"),
		})
		if err != nil {
			return err
		}
		return output(result, toRows(result.Feeds))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	for _, c := range []*cobra.Command{searchByTermCmd, searchByTitleCmd, searchByPersonCmd, searchMusicCmd} {
		c.Flags().StringP("query", "q", "", "Search query (required)")
		c.Flags().Int("max", 0, "Maximum number of results")
		c.Flags().Bool("fulltext", false, "Return full text description")
		searchCmd.AddCommand(c)
	}

	for _, c := range []*cobra.Command{searchByTermCmd, searchByTitleCmd, searchMusicCmd} {
		c.Flags().String("val", "", "Value block type filter (any/lightning/hive/webmonetization)")
		c.Flags().Bool("clean", false, "Only non-explicit feeds")
	}

	for _, c := range []*cobra.Command{searchByTermCmd, searchByTitleCmd} {
		c.Flags().Bool("similar", false, "Include similar matches")
	}

	for _, c := range []*cobra.Command{searchByTermCmd, searchMusicCmd} {
		c.Flags().Bool("aponly", false, "Only feeds with an iTunes ID")
	}
}


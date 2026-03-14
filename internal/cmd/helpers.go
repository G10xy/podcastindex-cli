package cmd

import (
	"os"

	"github.com/spf13/viper"

	"github.com/G10xy/podcastindex-cli/internal/client"
	"github.com/G10xy/podcastindex-cli/internal/formatter"
)

func newClient() *client.Client {
	return client.NewClient(viper.GetString("api_key"), viper.GetString("api_secret"))
}

func output(response interface{}, rows []formatter.TableRow) error {
	format := viper.GetString("output_format")
	f, err := formatter.New(format)
	if err != nil {
		return err
	}
	if format == "json" {
		return f.Format(os.Stdout, response)
	}
	return f.Format(os.Stdout, rows)
}

func toRows[T formatter.TableRow](items []T) []formatter.TableRow {
	rows := make([]formatter.TableRow, len(items))
	for i := range items {
		rows[i] = items[i]
	}
	return rows
}

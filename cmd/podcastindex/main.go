package main

import (
	"os"

	"github.com/G10xy/podcastindex-cli/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "hayasui",
		Short: "Discord bot to get anime/manga/character/people data.",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "bot",
		Short: "Run bot",
		RunE: func(*cobra.Command, []string) error {
			return bot()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const name = "poke"

const version = "0.0.1"

func main() {
	app := &cli.App{
		Name:        name,
		Usage:       name,
		Version:     version,
		Description: "A command line application for Pokémon.",
		Commands: []*cli.Command{
			{
				Name:  "search",
				Usage: "Search for Pokémon.",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Searching for Pokémon...", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

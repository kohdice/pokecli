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
		Description: "A command line application for Pokémon",
		Commands: []*cli.Command{
			{
				Name:        "number",
				Description: "Search for Pokémon by National Pokédex Number",
				Usage:       "Search for Pokémon by National Pokédex Number",
				UsageText:   "poke number [number]",
				HelpName:    "number",
				ArgsUsage:   "[number]",
				Action:      searchByNumber,
			},
			{
				Name:        "name",
				Description: "Search for Pokémon by name.",
				Usage:       "Search for Pokémon by name.",
				UsageText:   "poke name [name]",
				HelpName:    "name",
				ArgsUsage:   "[name]",
				Action:      searchByName,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

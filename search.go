package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func searchByNumber(cCtx *cli.Context) error {
	if cCtx.Args().First() == "" {
		return fmt.Errorf("national pokédex number is required")
	}

	arg := cCtx.Args().First()
	if _, err := strconv.Atoi(arg); err != nil {
		return fmt.Errorf("invalid argument: %s", arg)
	}

	table := generateTable()
	table.Render()

	return nil
}

func searchByName(cCtx *cli.Context) error {
	if cCtx.Args().First() == "" {
		return fmt.Errorf("pokémon's name is required")
	}

	table := generateTable()
	table.Render()

	return nil
}

func generateTable(pokemon *pokemon) *tablewriter.Table {
	data := [][]string{
		{"Status", "HP", strconv.Itoa(pokemon.HP)},
		{"Status", "Attack", strconv.Itoa(pokemon.Attack)},
		{"Status", "Defense", strconv.Itoa(pokemon.Defense)},
		{"Status", "SpecialAttack", strconv.Itoa(pokemon.SpecialAttack)},
		{"Status", "SpecialDefense", strconv.Itoa(pokemon.SpecialDefense)},
		{"Status", "Speed", strconv.Itoa(pokemon.Speed)},
		{"Status", "BaseTotal", strconv.Itoa(pokemon.BaseTotal)},
	}

	for _, v := range pokemon.Types {
		data = append(data, []string{"Types", fmt.Sprintf("Type %d", v.Slot), v.PokemonType.TypeName})
	}

	for _, v := range pokemon.Abilities {
		if v.IsHidden {
			data = append(data, []string{"Abilities", "Hidden Ability", v.PokemonAbility.AbilityName})
		} else {
			data = append(data, []string{"Abilities", fmt.Sprintf("Ability %d", v.Slot), v.PokemonAbility.AbilityName})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{fmt.Sprintf("No:%d", pokemon.NationalPokedexNumber), pokemon.Name, ""})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.AppendBulk(data)

	return table
}

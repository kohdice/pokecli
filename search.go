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

func generateTable() *tablewriter.Table {
	data := [][]string{
		{"Status", "HP", "45"},
		{"Status", "Attack", "49"},
		{"Status", "Defense", "49"},
		{"Status", "SpecialAttack", "65"},
		{"Status", "SpecialDefense", "65"},
		{"Status", "Speed", "45"},
		{"Status", "BaseTotal", "318"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{fmt.Sprintf("No:%d", 1), "フシギダネ", "Kanto"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.AppendBulk(data)

	return table
}

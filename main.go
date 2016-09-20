package main

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
)

type Team struct {
	Position       string
	Name           string
	Points         string
	Games          string
	Wins           string
	Draws          string
	Looses         string
	GoalsFor       string
	GoalsAgainst   string
	GoalDifference string
	Percentage     string
}

func main() {
	doc, err := goquery.NewDocument("http://globoesporte.globo.com/futebol/brasileirao-serie-b/")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var teams [20]Team

	doc.Find(".tabela-times strong.tabela-times-time-nome").Each(func(i int, s *goquery.Selection) {
		teams[i] = Team{Position: strconv.Itoa(i + 1), Name: s.Text()}
	})

	teamIndex := 0
	fieldIndex := 0
	doc.Find(".tabela-pontos tbody td:not(.tabela-pontos-ultimos-jogos)").Each(func(i int, s *goquery.Selection) {
		if i > 0 && i%9 == 0 {
			fieldIndex = 0
			teamIndex = teamIndex + 1
		}

		// could have used directly a map but I was interested in learn how reflect works in golang
		reflect.ValueOf(&teams[teamIndex]).Elem().Field(fieldIndex + 2).SetString(s.Text())

		fieldIndex = fieldIndex + 1
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pos", "Time", "P", "J", "V", "E", "D", "GP", "GC", "SG", "%"})

	for _, team := range teams {
		table.Append([]string{
			team.Position,
			team.Name,
			team.Points,
			team.Games,
			team.Wins,
			team.Draws,
			team.Looses,
			team.GoalsFor,
			team.GoalsAgainst,
			team.GoalDifference,
			team.Percentage,
		})
	}

	table.Render()
}

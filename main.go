package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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
		fmt.Println(teams[teamIndex])
		reflect.ValueOf(&teams[teamIndex]).Elem().Field(fieldIndex + 2).SetString(s.Text())

		fieldIndex = fieldIndex + 1
	})

	for _, team := range teams {
		fmt.Println(team)
	}
}

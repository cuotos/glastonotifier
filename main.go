package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	efestivalsURL = "https://www.efestivals.co.uk/festivals/glastonbury/2022/lineup.shtml"
)

type Band struct {
	Name  string
	State State
	Stage string
}

type State int

const (
	Rumoured State = iota
	StronglyRumoured
	TBC
	Confirmed
)

func (s State) String() string {
	switch s {
	case Rumoured:
		return "Rumoured"
	case StronglyRumoured:
		return "Strongly Rumoured"
	case TBC:
		return "TBC"
	case Confirmed:
		return "Confirmed"
	}

	return "unknown"
}

func main() {
	c := colly.NewCollector()

	bands := []Band{}

	c.OnHTML("#panel5 .band", func(h *colly.HTMLElement) {
		b := extractBandFromInputString(h.Text)
		bands = append(bands, b)
	})
	c.Visit(efestivalsURL)

	for _, b := range bands {
		if b.State == StronglyRumoured {
			fmt.Printf("%+v\n", b.Name)
		}
	}
}

func extractBandFromInputString(input string) Band {
	r := regexp.MustCompile(` (?P<Date>.+), (?P<Stage>.+)\((?P<State>.+)\) (?P<Name>.+)`)
	found := r.FindStringSubmatch(input)

	name := strings.Join(strings.Fields(found[4]), " ")
	stage := found[2]

	var state State

	switch found[3] {
	case "R":
		state = Rumoured
	case "SR":
		state = StronglyRumoured
	case "TBC":
		state = TBC
	case "C":
		state = Confirmed
	}

	return Band{
		Name:  name,
		State: state,
		Stage: stage,
	}
}

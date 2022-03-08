package main

import (
	"regexp"
	"strings"
)

type BandI interface {
	Store() error
}

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

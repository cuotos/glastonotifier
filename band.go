package main

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-redis/redis"
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
	UnknownState State = iota
	Rumoured
	StronglyRumoured
	TBC
	Confirmed
)

func (s State) String() string {
	switch s {
	case UnknownState:
		return "Unknown State"
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

type BandRepository interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}

type repository struct {
	Client redis.Cmdable
}

func NewRedisRepository(Client redis.Cmdable) BandRepository {
	return &repository{Client}
}

func (r *repository) Set(key string, value interface{}, exp time.Duration) error {

	return r.Client.Set(key, value, exp).Err()
}

func (r *repository) Get(key string) (string, error) {
	get := r.Client.Get(key)
	return get.Result()
}

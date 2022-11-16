package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSand(t *testing.T) {
	b := Band{
		Name:  "maddog",
		State: UnknownState,
		Stage: "unknown",
	}

	e := json.NewEncoder(os.Stdout)
	e.Encode(b)

}

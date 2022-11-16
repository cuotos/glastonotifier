package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBandFromHTMLElement(t *testing.T) {
	tcs := []struct {
		InputText    string
		ExpectedBand Band
	}{
		{
			" Wed 1st, Some Stage(R) A Band",
			Band{Name: "A Band", State: Rumoured, Stage: "Some Stage"},
		},
		{
			" Wed 1st, Some Stage(SR) A Band",
			Band{Name: "A Band", State: StronglyRumoured, Stage: "Some Stage"},
		},
		{
			" Fri 24th, Other Stage(C) Blossoms",
			Band{Name: "Blossoms", State: Confirmed, Stage: "Other Stage"},
		},
		{
			" day TBC, Lost Horizon Sauna Solar Stage(TBC) Papaphone",
			Band{Name: "Papaphone", State: TBC, Stage: "Lost Horizon Sauna Solar Stage"},
		},
		{
			" day TBC, unknown stage(TBC) Woody  Cook",
			Band{Name: "Woody Cook", State: TBC, Stage: "unknown stage"},
		},
	}

	for _, tc := range tcs {
		actual := extractBandFromInputString(tc.InputText)
		assert.Equal(t, tc.ExpectedBand, actual)
	}
}

func TestMarshalBand(t *testing.T) {
	tcs := []struct {
		InputBand    Band
		ExpectedJson string
	}{
		{
			InputBand: Band{
				Name:  "Dans Band",
				State: TBC,
				Stage: "A Stage",
			},
			ExpectedJson: `{"Name":"Dans Band","State":3,"Stage":"A Stage"}`,
		},
	}

	for _, tc := range tcs {
		actual, err := json.Marshal(tc.InputBand)
		require.NoError(t, err)

		assert.Equal(t, tc.ExpectedJson, string(actual))
	}
}

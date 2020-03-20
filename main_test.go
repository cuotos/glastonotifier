package main

import (
	"glastorumornotifier/testdata"
	"strings"
	"testing"
)

func TestCheckDifferences(t *testing.T) {
	raw1 := testdata.EventOne
	raw2 := testdata.EventTwo

	e1, _ := GetEventFromResponseBody(strings.NewReader(raw1))
	e2, _ := GetEventFromResponseBody(strings.NewReader(raw2))

	findDifferences(e1, e2)
}
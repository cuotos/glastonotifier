package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const (
	efestivalsURL = "https://www.efestivals.co.uk/festivals/glastonbury/2022/lineup.shtml"
)

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

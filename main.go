package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gocolly/colly/v2"
)

const (
	efestivalsURL = "https://www.efestivals.co.uk/festivals/glastonbury/2023/lineup.shtml"
)

var (
	ctx = context.Background()
)

func mustGet(envvar string) string {
	if val := os.Getenv(envvar); val != "" {
		return val
	}
	log.Fatalf("required env var missing %s", envvar)
	return ""
}

func main() {
	mustGet("REDIS_URL")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bg := BandRetriever{}
	bands := bg.GetAllBands()

	db := redis.NewClient(&redis.Options{Addr: mustGet("REDIS_URL")})

	if _, err := db.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("unable to connect to redis: %w", err)
	}

	repo := NewRedisRespository(db)
	// repo := FileRepo{}

	// makge sure the repo of choice implements Repository, not critical, just fun to do.
	var _ Repository = repo

	for _, b := range bands {
		key := strings.ReplaceAll(fmt.Sprintf("%s-%s", b.Name, b.Stage), " ", "_")
		repo.Store(ctx, key, b)
	}

	return nil
}

type BandRetriever struct{}

func (g *BandRetriever) GetAllBands() []Band {
	c := colly.NewCollector(colly.AllowURLRevisit())

	bands := []Band{}

	// Get the if the html panel holding the a-to-z list
	var panelId string
	c.OnHTML(`a:contains("A-to-Z")`, func(h *colly.HTMLElement) {
		panelId = h.Attr("href")
	})
	c.Visit(efestivalsURL)

	aToZPanelQuery := fmt.Sprintf("%s .band", panelId)

	// create a new collector that doesnt care about the a-to-z panel.
	// but everything else is the same.
	c = c.Clone()

	c.OnHTML(aToZPanelQuery, func(h *colly.HTMLElement) {
		b := extractBandFromInputString(h.Text)
		bands = append(bands, b)
	})
	c.Visit(efestivalsURL)

	return bands
}

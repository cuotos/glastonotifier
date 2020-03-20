package main

import (
	"fmt"
	"glastorumornotifier/songkick"
	"log"
	"os"
	"reflect"
	"time"
)

const (
	songkickAPIBasetURL = "https://api.songkick.com/api/3.0/events/"
)

var (
	songkickAPIKey string
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {

	songkickAPIKey = os.Getenv("SONGKICK_API_KEY")
	if songkickAPIKey == "" {
		log.Fatal("missing SONGKICK_API_KEY")
	}

	c := songkick.NewClient(songkickAPIBasetURL, songkickAPIKey)

	var cacheEvent *songkick.Event

	ticker := time.NewTicker(time.Second * 5)
	for ; true; <-ticker.C {
		e, err := c.GetEvent("39010381-glastonbury-festival-2020")
		if err != nil {
			log.Fatal(err)
		}

		if cacheEvent == nil {
			cacheEvent = e
		} else {
			findDifferences(cacheEvent, e)
		}
	}
}

func findDifferences(e1, e2 *songkick.Event) {

	for _, i := range e1.Performance {
		found := false
		for _, j := range e2.Performance {
			if reflect.DeepEqual(i, j) {
				found = true
			}
		}
		if !found {
			fmt.Println(i.DisplayName)
		}
	}
}
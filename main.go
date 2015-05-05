package main

import (
	"fmt"
	"os"

	"github.com/milk/hummingbird"
	"github.com/milk/nyaa"
)

func printErrors(errs []error) {
	for _, err := range errs {
		fmt.Println(err)
	}
}

func main() {
	username := os.Args[1]
	if len(username) == 0 {
		fmt.Printf("You must specify your HummingBird username.")
		return
	}

	api := hummingbird.NewAPI()
	errs, watching := api.Library(username, "currently-watching")
	if len(errs) != 0 {
		printErrors(errs)
		return
	}

	var animes []string

	for _, entry := range watching {
		animes = append(animes, entry.Anime.Title)
		if len(entry.Anime.Alternate_title) != 0 {
			animes = append(animes, entry.Anime.Alternate_title)
		}
	}

	nyaaApi := nyaa.NewAPI()

	for _, anime := range animes {
		fmt.Printf("Anime: %s\n", anime)
		entries, errs := nyaaApi.Search(anime)
		if len(errs) != 0 {
			printErrors(errs)
		}

		for _, entry := range entries {
			fmt.Printf("Name: %s - %s\nDownload: %s\n", entry.Name, entry.Link, entry.Torrent)
		}
		fmt.Printf("\n")
	}

}

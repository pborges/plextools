package main

import (
	"github.com/pborges/plextools"
	"fmt"
	"path"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage", os.Args[0], "<plex addr:port> [rename]")
		return
	}
	doRename := false

	c := plextools.Dial(os.Args[1])
	if len(os.Args) > 2 {
		if os.Args[2] == "rename" {
			doRename = true
			fmt.Println("I will actually rename the files")
		}
	}

	fmt.Println("Shows ------------------------------------------------------")
	shows, err := c.Shows()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, show := range shows {
		fmt.Println(show.Title)
		for _, episode := range show.Episodes {
			check(episode, doRename)
		}
	}

	fmt.Println("Movies -----------------------------------------------------")
	movies, err := c.Movies()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, movie := range movies {
		check(movie, doRename)
	}
}

func check(e plextools.NamedEntry, doRename bool) {
	ext := path.Ext(e.FilePath())
	diskPath := path.Dir(e.FilePath())
	newFilename := path.Join(diskPath, e.FormattedFileName()+ext)

	filenameMismatch := e.FilePath() != newFilename

	filenameCheck := "✓"
	if filenameMismatch {
		filenameCheck = " "
	}

	fmt.Printf("\t[%s] %s", filenameCheck, e.FilePath())

	if filenameMismatch {
		fmt.Printf(" -> %s", newFilename)
		if doRename {
			if err := os.Rename(e.FilePath(), newFilename); err != nil {
				fmt.Println()
				fmt.Println("ERROR:", err)
			}
		}
	}
	fmt.Println()
}

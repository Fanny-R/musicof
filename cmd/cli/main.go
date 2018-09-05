package main

import (
	"log"

	"github.com/jlevesy/musicof/core"
)

func main() {
	c := core.NewChooser()
	chosen, err := c.Choose()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(chosen)
}

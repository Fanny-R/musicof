package main

import (
	"log"

	"github.com/jlevesy/musicof/core"
)

func main() {
	choices := []string{"Pikachu", "Chuchmur", "Ronflex", "Caninos"}
	provider := core.NewDummyChoicesProvider(choices)
	c := core.NewChooser(provider)
	chosen, err := c.Choose()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(chosen)
}

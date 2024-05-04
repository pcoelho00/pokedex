package main

import (
	"fmt"

	"github.com/pcoelho00/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	fmt.Println("Starting loop:")
	startRepl(&cfg)
}

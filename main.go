package main

import (
	"fmt"
	"time"

	"github.com/pcoelho00/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	currentAreaName     *string
	nextLocationURL     *string
	previousLocationURL *string
	Pokedex             map[string]pokeapi.PokemonInfo
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		Pokedex:       make(map[string]pokeapi.PokemonInfo),
	}

	fmt.Println("Starting loop:")
	startRepl(&cfg)
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Available Commands:")
	commands := createNewCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting program")
	os.Exit(0)
	return nil
}

func commandMapNext(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous
	return nil
}

func commandMapPrevious(cfg *config, args ...string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas: ")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous
	return nil
}

func commandGetArea(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no location provided")
	}

	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Printf("Pokemon found in the area:\n")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", Title(pokemon.Pokemon.Name))
	}

	cfg.currentAreaName = &locationAreaName

	return nil
}

func commandPokemonStats(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]
	Pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return errors.New("pokemon not registered at pokedex")
	}

	fmt.Printf("\nPokemon: %s Id %v:\n", Title(Pokemon.Name), Pokemon.ID)
	fmt.Printf("Height: %v\n", Pokemon.Height)
	fmt.Printf("Weight: %v\n", Pokemon.Weight)
	fmt.Println("Stat Name: Base Value")
	for _, pokeinfo := range Pokemon.Stats {
		fmt.Printf(" - %s: %v\n", pokeinfo.Stat.Name, pokeinfo.BaseStat)
	}
	fmt.Println("Types:")
	for _, poke_type := range Pokemon.Types {
		fmt.Println(poke_type.Type.Name)
	}

	return nil
}

func commandCatchPokemon(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]

	if cfg.currentAreaName == nil {
		return errors.New("need to explore an area first")
	}

	pokemonEncounter, err := cfg.pokeapiClient.GetPokeInfoAtArea(*cfg.currentAreaName, pokemonName)
	if err != nil {
		return err
	}

	if pokemonEncounter.Pokemon.Name == "" {
		return fmt.Errorf("%s not found at current area %s", pokemonName, *cfg.currentAreaName)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonEncounter.Pokemon.Name)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.Pokedex[pokemonName] = pokemon

	return nil

}

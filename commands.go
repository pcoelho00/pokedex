package main

import (
	"errors"
	"fmt"
	"os"
)

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
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandPokemonStats(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]
	Pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon %s Id %v:\n", Pokemon.Name, Pokemon.ID)
	for _, pokeinfo := range Pokemon.Stats {
		fmt.Printf("%s Base: %v\n", pokeinfo.Stat.Name, pokeinfo.BaseStat)
	}

	return nil
}

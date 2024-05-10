package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func createNewCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the next Location Areas",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous Location Areas",
			callback:    commandMapPrevious,
		},
		"explore": {
			name:        "explore {area_name}",
			description: "List the pokemon in the Area",
			callback:    commandGetArea,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Shows the Pokemon Base Stats",
			callback:    commandPokemonStats,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Shows the Pokemon Base Stats",
			callback:    commandCatchPokemon,
		},
	}
}

func cleanInput(text string) []string {

	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		commands := createNewCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid command: ", commandName)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

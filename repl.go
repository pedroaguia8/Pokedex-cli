package main

import (
	"strings"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

type config struct {
	BaseUrl              *string
	NextLocationArea     *string
	PreviousLocationArea *string
	Cache                *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area>",
			description: "List of all the Pokémon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Throw a Pokeball at a pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(text)
	for i := 0; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}

	return result
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCliCommands() {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func commandMap(config *config) error {
	if config.Next == nil {
		fmt.Println("you're on the last page")
		return nil
	}

	mapRes, err := pokeapi.GetLocationAreas(*config.Next, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	config.Next = mapRes.Next
	config.Previous = mapRes.Previous

	for _, result := range mapRes.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	mapRes, err := pokeapi.GetLocationAreas(*config.Previous, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	config.Next = mapRes.Next
	config.Previous = mapRes.Previous

	for _, result := range mapRes.Results {
		fmt.Println(result.Name)
	}

	return nil
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
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(text)
	for i := 0; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}

	return result
}

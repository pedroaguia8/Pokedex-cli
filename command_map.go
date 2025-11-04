package main

import (
	"fmt"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandMap(config *config, _ []string) error {
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

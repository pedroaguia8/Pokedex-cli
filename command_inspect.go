package main

import (
	"fmt"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandInspect(config *config, commandParameters []string) error {
	if len(commandParameters) == 0 {
		return fmt.Errorf("this command takes a Pokemon name as parameter: inspect <pokemon>")
	}
	pokemonName := commandParameters[0]

	if _, ok := config.Pokedex[pokemonName]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fullUrl := *config.BaseUrl + "/pokemon/" + pokemonName
	pokemon, err := pokeapi.GetPokemon(fullUrl, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting pokemon: %w", err)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types :\n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokeType.Type.Name)
	}
	return nil
}

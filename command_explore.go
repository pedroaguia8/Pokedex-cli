package main

import (
	"fmt"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandExplore(config *config, commandParameters []string) error {
	if len(commandParameters) == 0 {
		return fmt.Errorf("this command takes an area name as parameter: explore <area>")
	}

	fullUrl := *config.BaseUrl + "/location-area/" + commandParameters[0]
	pokemonRes, err := pokeapi.GetAreaPokemon(fullUrl, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting area pokemon: %w", err)
	}

	for _, pokemon := range pokemonRes.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

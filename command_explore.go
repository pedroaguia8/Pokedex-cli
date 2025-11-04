package main

import (
	"fmt"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandExplore(config *config, commandParameters []string) error {
	if len(commandParameters) == 0 {
		return fmt.Errorf("this command takes an area name as parameter: explore <area>")
	}
	areaToSearch := commandParameters[0]

	fullUrl := *config.BaseUrl + "/location-area/" + areaToSearch
	pokemonRes, err := pokeapi.GetAreaPokemon(fullUrl, config.Cache)
	if err != nil {
		return fmt.Errorf("error getting pokemon from area: %w", err)
	}

	fmt.Println("Exploring " + areaToSearch + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonRes.PokemonEncounters {
		fmt.Println(" -" + pokemon.Pokemon.Name)
	}
	return nil
}

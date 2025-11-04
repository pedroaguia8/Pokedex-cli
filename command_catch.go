package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
)

func commandCatch(config *config, commandParameters []string) error {
	if len(commandParameters) == 0 {
		return fmt.Errorf("this command takes a Pokemon name as parameter: catch <pokemon>")
	}
	pokemonToCatch := commandParameters[0]

	fullUrl := *config.BaseUrl + "/pokemon/" + pokemonToCatch
	pokemon, err := pokeapi.GetPokemon(fullUrl, config.Cache)
	if err != nil {
		return fmt.Errorf("error retrieving pokemon data: %w", err)
	}

	if err := throwPokeball(pokemon); err != nil {
		return fmt.Errorf("error throwing pokeball: %w", err)
	}
	return nil
}

func throwPokeball(pokemon pokeapi.PokemonResponse) error {
	if pokemon.BaseExperience == nil {
		return fmt.Errorf("error retrieving pokemon base experience")
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	exp := *pokemon.BaseExperience
	// with this formula the lowest experience Pokemon (Snom) will have 100% chance of getting caught
	// and the highest one (Blissey) will have 1/6 chance of getting caught
	prob := 113.8 / (float64(exp) + 74.8)
	if prob >= rand.Float64() {
		fmt.Println(pokemon.Name + " was caught!")
		return nil
	}
	fmt.Println(pokemon.Name + " escaped!")
	return nil
}

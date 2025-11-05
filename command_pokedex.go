package main

import "fmt"

func commandPokedex(config *config, par []string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your Pokédex is empty. Go catch some Pokémon!")
		return nil
	}

	for _, pokemon := range config.Pokedex {
		fmt.Println(" - " + pokemon.Name)
	}
	return nil
}

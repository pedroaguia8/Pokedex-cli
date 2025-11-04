package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

func GetPokemon(url string, cache *pokecache.Cache) (PokemonResponse, error) {
	if cachedData, ok := cache.Get(url); ok {
		pokemon := PokemonResponse{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonResponse{}, fmt.Errorf("error unmarshalling cached pokemon: %w", err)
		}
		log.Println("Using cached search")
		return pokemon, nil
	}

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error making request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error closing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return PokemonResponse{},
			fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	pokemonRes := PokemonResponse{}
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	cache.Add(url, body)

	return pokemonRes, nil
}

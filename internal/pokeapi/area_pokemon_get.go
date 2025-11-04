package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

func GetAreaPokemon(url string, cache *pokecache.Cache) (AreaPokemonsResponse, error) {
	if cachedData, ok := cache.Get(url); ok {
		pokemon := AreaPokemonsResponse{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return AreaPokemonsResponse{}, fmt.Errorf("error unmarshalling cached area's pokemon: %w", err)
		}
		log.Println("Using cached search")
		return pokemon, nil
	}

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return AreaPokemonsResponse{}, fmt.Errorf("error making request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return AreaPokemonsResponse{}, fmt.Errorf("error closing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return AreaPokemonsResponse{},
			fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	pokemonRes := AreaPokemonsResponse{}
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return AreaPokemonsResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	cache.Add(url, body)

	return pokemonRes, nil
}

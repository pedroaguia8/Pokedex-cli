package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) (LocationAreaResponse, error) {
	if cachedData, ok := cache.Get(url); ok {
		areas := LocationAreaResponse{}
		err := json.Unmarshal(cachedData, &areas)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error unmarshalling cached areas: %w", err)
		}
		log.Println("Using cached search")
		return areas, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error making request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error closing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return LocationAreaResponse{},
			fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	mapRes := LocationAreaResponse{}
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	cache.Add(url, body)

	return mapRes, nil
}

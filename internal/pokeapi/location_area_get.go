package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetLocationAreas(url string) (LocationAreaResponse, error) {
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
		return LocationAreaResponse{}, fmt.Errorf("error unmarshalling responde: %w", err)
	}

	return mapRes, nil
}

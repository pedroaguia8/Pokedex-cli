package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MapResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (MapResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return MapResponse{}, fmt.Errorf("error making request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return MapResponse{}, fmt.Errorf("error closing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return MapResponse{},
			fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	mapRes := MapResponse{}
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return MapResponse{}, fmt.Errorf("error unmarshalling responde: %w", err)
	}

	return mapRes, nil
}

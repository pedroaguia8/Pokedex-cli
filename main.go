package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	config := config{
		Next: "https://pokeapi.co/api/v2/location-area/",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		command, ok := getCliCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCliCommands() {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

func commandMap(config *config) error {
	res, err := http.Get(config.Next)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return fmt.Errorf("error closing response body: %w", err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	mapRes := mapResponse{}
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return fmt.Errorf("error unmarshalling responde: %w", err)
	}

	if mapRes.Next != nil {
		config.Next = *mapRes.Next
	}
	if mapRes.Previous != nil {
		config.Previous = *mapRes.Previous
	}

	for _, result := range mapRes.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(config *config) error {

	return nil
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(text)
	for i := 0; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}

	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

type mapResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

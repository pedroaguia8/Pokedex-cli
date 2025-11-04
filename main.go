package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokeapi"
	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)

	baseUrl := "https://pokeapi.co/api/v2"
	nextLocationAreaUrl := baseUrl + "/location-area/?offset=0&limit=20"
	config := config{
		Cache:            cache,
		BaseUrl:          &baseUrl,
		NextLocationArea: &nextLocationAreaUrl,
		Pokedex:          map[string]pokeapi.Pokemon{},
	}

	fmt.Println("Welcome to the Pokedex!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		commandParameters := input[1:]
		command, ok := getCliCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(&config, commandParameters)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

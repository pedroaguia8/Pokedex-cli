package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/pedroaguia8/Pokedex-cli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)

	initialURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	config := config{
		Cache: cache,
		Next:  &initialURL,
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

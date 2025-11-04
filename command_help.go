package main

import "fmt"

func commandHelp(*config, []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCliCommands() {
		fmt.Println(command.name + ": " + command.description)
	}
	return nil
}

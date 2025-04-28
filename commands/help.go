package commands

import (
	"fmt"
)

// Help displays a help message with all available commands
func Help(cfg *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range GetCommands() {
		fmt.Println(key + ": " + value.Description)
	}
	return nil
}

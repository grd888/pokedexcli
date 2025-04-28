package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grd888/pokedexcli/commands"
	"github.com/grd888/pokedexcli/internal/pokecache"
)

func main() {
	// Initialize the cache
	cache := pokecache.NewCache()
	
	// Initialize commands
	commandMap := commands.InitCommands(cache)
	
	// Create a config to track pagination state
	var config commands.Config

	// Start the CLI loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := cleanInput(scanner.Text())

		if len(line) == 0 {
			continue
		}
		
		command := line[0]
		args := []string{}
		if len(line) > 1 {
			args = line[1:]
		}
		
		handler, exists := commandMap[command]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		
		if handler.Callback != nil {
			err := handler.Callback(&config, args)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

// cleanInput cleans and tokenizes the input string
func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}

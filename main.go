package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/grd888/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next     string
	Previous string
}

var locationAreaURL string = "https://pokeapi.co/api/v2/location-area"
var commands map[string]cliCommand
var conf config
var pokeCache *pokecache.Cache

func main() {
	pokeCache = pokecache.NewCache()

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
	}

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
		
		handler, exists := commands[command]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		
		if handler.callback != nil {
			err := handler.callback(&conf, args)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text) // Use strings.Fields
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Println(key + ": " + value.description)
	}
	return nil
}
func commandMap(cfg *config, args []string) error {
	url := locationAreaURL
	if cfg.Next != "" {
		url = cfg.Next
	}

	// Check cache first
	if data, ok := pokeCache.Get(url); ok {
		var locationAreaResponse LocationAreaResponse
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return err
		}
		cfg.Next = locationAreaResponse.Next
		cfg.Previous = locationAreaResponse.Previous

		fmt.Println("Location Areas (from cache):")
		for _, locationArea := range locationAreaResponse.Results {
			fmt.Println("-", locationArea.Name)
		}
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Cache the response
	pokeCache.Add(url, body)

	var locationAreaResponse LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return err
	}

	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous
	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}

func commandMapB(cfg *config, args []string) error {
	var url string
	if cfg.Previous != "" {
		url = cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var locationAreaResponse LocationAreaResponse
	err = json.NewDecoder(resp.Body).Decode(&locationAreaResponse)
	if err != nil {
		return err
	}
	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}
	
	locationAreaName := args[0]
	
	// Construct the URL for the specific location area
	url := fmt.Sprintf("%s/%s", locationAreaURL, locationAreaName)
	
	// Check cache first
	if data, ok := pokeCache.Get(url); ok {
		return processLocationAreaData(data)
	}
	
	fmt.Printf("Exploring %s...\n", locationAreaName)
	
	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("location area '%s' not found", locationAreaName)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	// Cache the response
	pokeCache.Add(url, body)
	
	return processLocationAreaData(body)
}

func getCommands() map[string]cliCommand {
	return commands
}

func processLocationAreaData(data []byte) error {
	type LocationAreaDetails struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		PokemonEncounters []struct {
			Pokemon struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}
	
	var locationArea LocationAreaDetails
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return err
	}
	
	fmt.Println("Found Pokemon:")
	
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println(" - No Pokemon found in this area.")
		return nil
	}
	
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	
	return nil
}

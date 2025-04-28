package commands

import (
	"github.com/grd888/pokedexcli/internal/pokecache"
)

// CLICommand represents a command that can be executed in the CLI
type CLICommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

// Config holds the state of the CLI application
type Config struct {
	Next     string
	Previous string
}

// LocationAreaResponse represents the response from the location area API
type LocationAreaResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// LocationArea represents a location area in the Pokemon world
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreaDetails represents the detailed information about a location area
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

var (
	// LocationAreaURL is the base URL for the location area API
	LocationAreaURL = "https://pokeapi.co/api/v2/location-area"
	
	// PokeCache is the cache for API responses
	PokeCache *pokecache.Cache
	
	// Commands is a map of all available commands
	Commands map[string]CLICommand
)

// InitCommands initializes the commands map
func InitCommands(cache *pokecache.Cache) map[string]CLICommand {
	PokeCache = cache
	
	Commands = map[string]CLICommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"map": {
			Name:        "map",
			Description: "Show location areas",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show previous location areas",
			Callback:    MapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area",
			Callback:    Explore,
		},
	}
	
	return Commands
}

// GetCommands returns the commands map
func GetCommands() map[string]CLICommand {
	return Commands
}

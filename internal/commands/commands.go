package commands

import (
	"github.com/grd888/pokedexcli/internal/models"
	"github.com/grd888/pokedexcli/internal/pokecache"
)

var (
	// LocationAreaURL is the base URL for the location area API
	LocationAreaURL = "https://pokeapi.co/api/v2/location-area"
	
	// Cache is the cache for API responses
	Cache *pokecache.Cache
	
	// CommandMap is a map of all available commands
	CommandMap map[string]models.Command
)

// Initialize initializes the commands package
func Initialize(cache *pokecache.Cache) map[string]models.Command {
	Cache = cache
	
	CommandMap = map[string]models.Command{
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
	
	return CommandMap
}

// GetCommands returns the commands map
func GetCommands() map[string]models.Command {
	return CommandMap
}

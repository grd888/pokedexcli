package commands

import (
	"github.com/grd888/pokedexcli/internal/api"
	"github.com/grd888/pokedexcli/internal/models"
	"github.com/grd888/pokedexcli/internal/pokecache"
)

var (
	// APIClient is the client for the PokeAPI
	APIClient *api.Client
	
	// CommandMap is a map of all available commands
	CommandMap map[string]models.Command
)

// Initialize initializes the commands package
func Initialize(cache *pokecache.Cache) map[string]models.Command {
	APIClient = api.NewClient(cache)
	
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
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon",
			Callback:    Catch,
		},
	}
	
	return CommandMap
}

// GetCommands returns the commands map
func GetCommands() map[string]models.Command {
	return CommandMap
}

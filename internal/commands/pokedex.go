package commands

import (
	"fmt"
	"sort"

	"github.com/grd888/pokedexcli/internal/models"
)

// Pokedex displays a list of all caught Pokemon
func Pokedex(cfg *models.Config, args []string) error {
	// Check if any Pokemon have been caught
	if len(CaughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon!")
		return nil
	}

	// Get all Pokemon names
	var pokemonNames []string
	for name := range CaughtPokemon {
		pokemonNames = append(pokemonNames, name)
	}

	// Sort the names alphabetically for a consistent display
	sort.Strings(pokemonNames)

	// Display the list of caught Pokemon
	fmt.Println("Your Pokedex:")
	for _, name := range pokemonNames {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

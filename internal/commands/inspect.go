package commands

import (
	"fmt"
	"strings"

	"github.com/grd888/pokedexcli/internal/models"
)

// Inspect displays detailed information about a caught Pokemon
func Inspect(cfg *models.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	// Get the Pokemon name from args
	pokemonName := strings.ToLower(args[0])

	// Check if the Pokemon has been caught
	caughtPokemon, ok := CaughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Display Pokemon details
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", caughtPokemon.Height)
	fmt.Printf("Weight: %d\n", caughtPokemon.Weight)
	
	// Display stats
	fmt.Println("Stats:")
	for _, stat := range caughtPokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	
	// Display types
	fmt.Println("Types:")
	for _, t := range caughtPokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

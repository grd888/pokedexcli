package commands

import (
	"fmt"

	"github.com/grd888/pokedexcli/internal/models"
)

// Explore displays Pokemon in a specific location area
func Explore(cfg *models.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}
	
	locationAreaName := args[0]
	
	fmt.Printf("Exploring %s...\n", locationAreaName)
	
	// Get location area details from the API
	locationAreaDetails, err := APIClient.GetLocationAreaDetails(locationAreaName)
	if err != nil {
		return err
	}
	
	return displayLocationAreaDetails(locationAreaDetails)
}

// displayLocationAreaDetails displays the Pokemon in a location area
func displayLocationAreaDetails(locationArea models.LocationAreaDetails) error {
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

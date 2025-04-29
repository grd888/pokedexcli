package commands

import (
	"fmt"

	"github.com/grd888/pokedexcli/internal/models"
)

// Map displays the location areas from the PokeAPI
func Map(cfg *models.Config, args []string) error {
	pageURL := ""
	if cfg.Next != "" {
		pageURL = cfg.Next
	}

	// Get location areas from the API
	locationAreaResponse, err := APIClient.GetLocationAreas(pageURL)
	if err != nil {
		return err
	}

	// Update config with pagination URLs
	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	// Display location areas
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}

// MapB displays the previous page of location areas
func MapB(cfg *models.Config, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	// Get previous page of location areas from the API
	locationAreaResponse, err := APIClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	// Update config with pagination URLs
	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	// Display location areas
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}

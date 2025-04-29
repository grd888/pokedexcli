package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grd888/pokedexcli/internal/models"
)

// Explore displays Pokemon in a specific location area
func Explore(cfg *models.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}
	
	locationAreaName := args[0]
	
	// Construct the URL for the specific location area
	url := fmt.Sprintf("%s/%s", LocationAreaURL, locationAreaName)
	
	// Check cache first
	if data, ok := Cache.Get(url); ok {
		fmt.Printf("Exploring %s (from cache)...\n", locationAreaName)
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
	Cache.Add(url, body)
	
	return processLocationAreaData(body)
}

// processLocationAreaData processes the location area data and displays Pokemon
func processLocationAreaData(data []byte) error {
	var locationArea models.LocationAreaDetails
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

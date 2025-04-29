package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grd888/pokedexcli/internal/models"
)

// Map displays the location areas from the PokeAPI
func Map(cfg *models.Config, args []string) error {
	url := LocationAreaURL
	if cfg.Next != "" {
		url = cfg.Next
	}

	// Check cache first
	if data, ok := Cache.Get(url); ok {
		var locationAreaResponse models.LocationAreaResponse
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
	Cache.Add(url, body)

	var locationAreaResponse models.LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
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

// MapB displays the previous page of location areas
func MapB(cfg *models.Config, args []string) error {
	var url string
	if cfg.Previous != "" {
		url = cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	
	// Check cache first
	if data, ok := Cache.Get(url); ok {
		var locationAreaResponse models.LocationAreaResponse
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
	Cache.Add(url, body)

	var locationAreaResponse models.LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
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

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/grd888/pokedexcli/internal/models"
)

const (
	// PokemonEndpoint is the endpoint for Pokemon data
	PokemonEndpoint = BaseURL + "/pokemon"
)

// GetPokemon retrieves a Pokemon by name from the PokeAPI
func (c *Client) GetPokemon(name string) (models.Pokemon, error) {
	// Normalize the name (lowercase)
	name = strings.ToLower(name)
	
	// Construct the URL
	url := fmt.Sprintf("%s/%s", PokemonEndpoint, name)

	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		var pokemon models.Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return models.Pokemon{}, err
		}
		return pokemon, nil
	}

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		return models.Pokemon{}, err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return models.Pokemon{}, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Pokemon{}, err
	}

	// Store in cache
	c.cache.Add(url, body)

	// Unmarshal the JSON
	var pokemon models.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return models.Pokemon{}, err
	}

	return pokemon, nil
}

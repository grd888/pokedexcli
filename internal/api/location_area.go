package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grd888/pokedexcli/internal/models"
)

// GetLocationAreas fetches a list of location areas
func (c *Client) GetLocationAreas(pageURL string) (models.LocationAreaResponse, error) {
	url := LocationAreaEndpoint
	if pageURL != "" {
		url = pageURL
	}

	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		var locationAreaResponse models.LocationAreaResponse
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return models.LocationAreaResponse{}, err
		}
		return locationAreaResponse, nil
	}

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		return models.LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.LocationAreaResponse{}, err
	}

	// Cache the response
	c.cache.Add(url, body)

	// Parse the response
	var locationAreaResponse models.LocationAreaResponse
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return models.LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}

// GetLocationAreaDetails fetches details for a specific location area
func (c *Client) GetLocationAreaDetails(name string) (models.LocationAreaDetails, error) {
	url := fmt.Sprintf("%s/%s", LocationAreaEndpoint, name)

	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		var locationAreaDetails models.LocationAreaDetails
		err := json.Unmarshal(data, &locationAreaDetails)
		if err != nil {
			return models.LocationAreaDetails{}, err
		}
		return locationAreaDetails, nil
	}

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		return models.LocationAreaDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.LocationAreaDetails{}, fmt.Errorf("location area '%s' not found", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.LocationAreaDetails{}, err
	}

	// Cache the response
	c.cache.Add(url, body)

	// Parse the response
	var locationAreaDetails models.LocationAreaDetails
	err = json.Unmarshal(body, &locationAreaDetails)
	if err != nil {
		return models.LocationAreaDetails{}, err
	}

	return locationAreaDetails, nil
}

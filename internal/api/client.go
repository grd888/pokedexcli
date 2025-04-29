package api

import (
	"github.com/grd888/pokedexcli/internal/pokecache"
)

const (
	// BaseURL is the base URL for the PokeAPI
	BaseURL = "https://pokeapi.co/api/v2"
	
	// LocationAreaEndpoint is the endpoint for location areas
	LocationAreaEndpoint = BaseURL + "/location-area"
)

// Client is a client for the PokeAPI
type Client struct {
	cache *pokecache.Cache
}

// NewClient creates a new PokeAPI client
func NewClient(cache *pokecache.Cache) *Client {
	return &Client{
		cache: cache,
	}
}

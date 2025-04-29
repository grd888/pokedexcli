package models

// LocationAreaResponse represents the response from the location area API
type LocationAreaResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// LocationArea represents a location area in the Pokemon world
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreaDetails represents the detailed information about a location area
type LocationAreaDetails struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

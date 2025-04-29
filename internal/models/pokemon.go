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

// Pokemon represents a Pokemon from the PokeAPI
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Types          []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

// CaughtPokemon represents a Pokemon that has been caught by the player
type CaughtPokemon struct {
	Pokemon
	CaughtAt string `json:"caught_at"`
}

package commands

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "strings"
    "time"

    "github.com/grd888/pokedexcli/internal/models"
)

// TryCapture attempts to capture a Pokémon based on its base_experience.
// Returns true if the capture is successful, false otherwise.
// If seed is not nil, it will be used to seed the random number generator for deterministic results (useful for testing).
func TryCapture(baseExperience int, seed *int64) bool {
    const maxBaseExperience = 608 // Max base_experience (e.g., Arceus)
    const maxProbability = 0.70   // Max capture probability (70% for weak Pokémon)
    const minProbability = 0.05   // Min capture probability (5% for strong Pokémon)

    // Validate input
    if baseExperience < 0 {
        baseExperience = 0 // Treat invalid input as minimum
    }

    // Calculate capture probability - using a more aggressive curve to make stronger Pokémon harder to catch
    // We're using a power of 0.8 to make the curve steeper for higher base experience values
    expFactor := math.Pow(float64(baseExperience)/float64(maxBaseExperience), 0.8)
    probability := maxProbability - expFactor * (maxProbability - minProbability)
    
    // Ensure probability doesn't go below minimum
    if probability < minProbability {
        probability = minProbability
    }

    // Generate random number between 0 and 1
    var randomValue float64
    if seed != nil {
        // Use seeded random source for deterministic results (testing)
        r := rand.New(rand.NewSource(*seed))
        randomValue = r.Float64()
    } else {
        // Use default random source for normal gameplay
        // Initialize the global random source if it hasn't been initialized yet
        // This ensures we get different results each time the game is played
        randomValue = rand.Float64()
    }

    // Return true if random value is less than capture probability
    return randomValue < probability
}

// CaughtPokemon is a map to store caught Pokemon
var CaughtPokemon = make(map[string]models.CaughtPokemon)

// Catch attempts to catch a Pokemon by name
func Catch(cfg *models.Config, args []string) error {
    if len(args) == 0 {
        return errors.New("you must provide a Pokemon name")
    }

    // Get the Pokemon name from args
    pokemonName := strings.ToLower(args[0])

    // Check if we've already caught this Pokemon
    if _, ok := CaughtPokemon[pokemonName]; ok {
        return fmt.Errorf("you've already caught %s", pokemonName)
    }

    fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
    
    // Add a small delay to build suspense
    time.Sleep(time.Millisecond * 500)
    
    // Fetch Pokemon data from the API
    pokemon, err := APIClient.GetPokemon(pokemonName)
    if err != nil {
        return fmt.Errorf("failed to find Pokemon '%s': %w", pokemonName, err)
    }

    // Show difficulty information based on base experience
    var difficultyText string
    if pokemon.BaseExperience < 100 {
        difficultyText = "It looks easy to catch!"
    } else if pokemon.BaseExperience < 200 {
        difficultyText = "It might be a challenge to catch."
    } else if pokemon.BaseExperience < 300 {
        difficultyText = "This one looks tough to catch!"
    } else {
        difficultyText = "This is a very powerful Pokémon! It will be extremely difficult to catch!"
    }
    
    fmt.Printf("\nYou found a level %d %s! %s\n", pokemon.BaseExperience/10+1, strings.Title(pokemonName), difficultyText)
    
    // Add another small delay for suspense
    time.Sleep(time.Millisecond * 500)
    
    // Show the Poké Ball animation
    fmt.Print("The Poké Ball wobbles...")
    time.Sleep(time.Millisecond * 300)
    
    // Attempt to capture the Pokemon
    if TryCapture(pokemon.BaseExperience, nil) {
        // Success! Add to caught Pokemon
        caughtPokemon := models.CaughtPokemon{
            Pokemon:  pokemon,
            CaughtAt: time.Now().Format(time.RFC3339),
        }
        CaughtPokemon[pokemonName] = caughtPokemon

        fmt.Printf("\n%s was caught!\n", pokemonName)
        fmt.Printf("You may now inspect it with the inspect command.\n")
        fmt.Println()
        return nil
    }

    // Failed to capture
    fmt.Printf("\n%s escaped!\n", pokemon.Name)
    fmt.Println("Maybe try again or catch a different Pokémon?")
    return nil
}
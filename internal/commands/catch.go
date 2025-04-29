package commands

import (
    "math/rand"
)

// TryCapture attempts to capture a Pokémon based on its base_experience.
// Returns true if the capture is successful, false otherwise.
// If seed is not nil, it will be used to seed the random number generator for deterministic results (useful for testing).
func TryCapture(baseExperience int, seed *int64) bool {
    const maxBaseExperience = 608 // Max base_experience (e.g., Arceus)
    const maxProbability = 0.95   // Max capture probability (95% for weak Pokémon)
    const minProbability = 0.1    // Min capture probability (10% for strong Pokémon)

    // Validate input
    if baseExperience < 0 {
        baseExperience = 0 // Treat invalid input as minimum
    }

    // Calculate capture probability
    probability := maxProbability - float64(baseExperience)/float64(maxBaseExperience) * (maxProbability - minProbability)
    
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
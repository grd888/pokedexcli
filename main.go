package main

import (
	"github.com/grd888/pokedexcli/internal/cli"
)

func main() {
	// Create a new CLI instance
	pokedex := cli.NewCLI()
	
	// Run the CLI application
	pokedex.Run()
}

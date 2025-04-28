package commands

import (
	"fmt"
	"os"
)

// Exit closes the Pokedex application
func Exit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

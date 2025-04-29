package commands

import (
	"fmt"
	"os"

	"github.com/grd888/pokedexcli/internal/models"
)

// Exit closes the Pokedex application
func Exit(cfg *models.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

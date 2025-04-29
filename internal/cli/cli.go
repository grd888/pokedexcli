package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grd888/pokedexcli/internal/api"
	"github.com/grd888/pokedexcli/internal/commands"
	"github.com/grd888/pokedexcli/internal/models"
	"github.com/grd888/pokedexcli/internal/pokecache"
)

// CLI represents the Pokedex CLI application
type CLI struct {
	cache      *pokecache.Cache
	apiClient  *api.Client
	config     *models.Config
	commands   map[string]models.Command
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	cache := pokecache.NewCache()
	apiClient := api.NewClient(cache)
	
	cli := &CLI{
		cache:     cache,
		apiClient: apiClient,
		config:    models.NewConfig(),
		commands:  make(map[string]models.Command),
	}
	
	// Register commands
	cli.registerCommands()
	
	return cli
}

// Run starts the CLI application
func (c *CLI) Run() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for a list of commands")
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := c.cleanInput(scanner.Text())
		
		if len(input) == 0 {
			continue
		}
		
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		
		command, exists := c.commands[commandName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		
		err := command.Callback(c.config, args)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// cleanInput cleans and tokenizes the input string
func (c *CLI) cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}

// registerCommands registers all available commands
func (c *CLI) registerCommands() {
	c.commands = commands.Initialize(c.cache)
}

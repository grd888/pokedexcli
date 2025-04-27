package main

import (
	"fmt"
	"strings"
  "bufio"
  "os"
  "net/http"
  "encoding/json"
)

type cliCommand struct {
  name string
  description string
  callback func(*config) error
}

type config struct {
  Next string
  Previous string
}

var locationAreaURL string = "https://pokeapi.co/api/v2/location-area"
var commands map[string]cliCommand
var conf config

func main() {
  
  commands = map[string]cliCommand{
    "help": {
      name: "help", 
      description: "Displays a help message", 
      callback: commandHelp,
    },
    "exit": {
      name: "exit", 
      description: "Exit the Pokedex", 
      callback: commandExit,
    },
    "map": {
      name: "map", 
      description: "Show location areas", 
      callback: commandMap,
    },
    "mapb": {
      name: "mapb", 
      description: "Show previous location areas", 
      callback: commandMapB,
    },
  }

  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    line := cleanInput(scanner.Text())

    command := line[0]
    handler := commands[command]
    if handler.callback != nil {
      err := handler.callback(&conf)
      if err != nil {
        fmt.Println("Error:", err)
      }
    } else {
      fmt.Println("Unknown command. Type 'help' for a list of commands.")
    }
  }
}

func cleanInput(text string) []string {
  text = strings.TrimSpace(text)
  text = strings.ToLower(text)
  return strings.Fields(text) // Use strings.Fields
}

func commandExit(cfg *config) error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}

func commandHelp(cfg *config) error {
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  fmt.Println()
  for key, value := range getCommands() {
    fmt.Println(key + ": " + value.description)
  }
  return nil
}
func commandMap(cfg *config) error {
  url := locationAreaURL
  if cfg.Next != "" {
    url = cfg.Next
  }
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  var locationAreaResponse LocationAreaResponse
  err = json.NewDecoder(resp.Body).Decode(&locationAreaResponse)
  if err != nil {
    return err
  }
  cfg.Next = locationAreaResponse.Next
  cfg.Previous = locationAreaResponse.Previous
  for _, locationArea := range locationAreaResponse.Results {
    fmt.Println(locationArea.Name)
  }
  fmt.Println()

  return nil
}

func commandMapB(cfg *config) error {
  url := locationAreaURL
  if cfg.Previous != "" {
    url = cfg.Previous
  } else {
    fmt.Println("you're on the first page")
    return nil
  }
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  var locationAreaResponse LocationAreaResponse
  err = json.NewDecoder(resp.Body).Decode(&locationAreaResponse)
  if err != nil {
    return err
  }
  cfg.Next = locationAreaResponse.Next
  cfg.Previous = locationAreaResponse.Previous
  for _, locationArea := range locationAreaResponse.Results {
    fmt.Println(locationArea.Name)
  }
  fmt.Println()

  return nil
}

func getCommands() map[string]cliCommand {
  return commands
}
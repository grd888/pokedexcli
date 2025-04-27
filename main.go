package main

import (
	"fmt"
	"strings"
  "bufio"
  "os"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

var commands map[string]cliCommand

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
  }

  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    line := cleanInput(scanner.Text())

    command := line[0]
    handler := commands[command]
    if handler.callback != nil {
      err := handler.callback()
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

func commandExit() error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
}

func commandHelp() error {
  fmt.Println("Welcome to the Pokedex!")
  fmt.Println("Usage:")
  fmt.Println()
  for key, value := range getCommands() {
    fmt.Println(key + ": " + value.description)
  }
  return nil
}

func getCommands() map[string]cliCommand {
  return commands
}
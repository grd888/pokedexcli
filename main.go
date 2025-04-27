package main

import (
	"fmt"
	"strings"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    line := cleanInput(scanner.Text())
    // get the first word from line
    command := line[0]
    fmt.Printf("Your command was: %s\n", command)
  }
}

func cleanInput(text string) []string {
  text = strings.TrimSpace(text)
  text = strings.ToLower(text)
  return strings.Fields(text) // Use strings.Fields
}
#!/bin/bash
# Test script for the pokedex command

# Build the application
go build -o pokedexcli

# Run the application with a sequence of commands
printf "pokedex\ncatch pidgey\ncatch caterpie\npokedex\nexit\n" | ./pokedexcli

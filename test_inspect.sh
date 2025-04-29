#!/bin/bash
# Test script for the inspect command

# Build the application
go build -o pokedexcli

# Run the application with a sequence of commands
printf "inspect pidgey\ncatch pidgey\ninspect pidgey\nexit\n" | ./pokedexcli

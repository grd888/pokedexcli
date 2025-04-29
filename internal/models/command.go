package models

// Command represents a command that can be executed in the CLI
type Command struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

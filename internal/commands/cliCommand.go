package commands

type cliCommand struct {
	name string
	description string
	Callback func(cliConfig *Config) error
}
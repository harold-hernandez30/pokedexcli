package commands

func GetCliCommand(command string) (cliCommand, bool) {
	cliCommand, ok := commandsMap[command]
	return cliCommand, ok
}

var commandsMap = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		Callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		Callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Fetches the next list of poke location areas",
		Callback: mapCommand,
	},
	"mapb": {
		name: "mapb",
		description: "Fetches the previous list of poke location areas",
		Callback: mapbCommand,
	},
}


package commands


var registry = map[string]cliCommand{}


func RegisterCommand(cmd string) {
	commandCli, ok := GetCliCommand(cmd)
	if ok {
		registry[cmd] = commandCli
	}
}

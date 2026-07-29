package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "show the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "show the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

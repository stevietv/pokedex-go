package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stevietv/pokedexgo/internal/pokeapi"
	"github.com/stevietv/pokedexgo/internal/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocations *string
	prevLocations *string
	pokecache     *pokecache.Cache
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			fmt.Println()
		}

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, ok := getCommands()[input[0]]
		parameters := strings.Join(input[1:], " ")

		if ok {
			err := command.callback(cfg, parameters)
			if err != nil {
				fmt.Printf("Error running command '%s': %v", command.name, err)
				fmt.Println()
			}
		} else {
			fmt.Println("Unknown Command")
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	return strings.Fields(text)
}

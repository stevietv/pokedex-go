package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
		}

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, ok := getCommands()[input[0]]

		if ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error running command '%s', %v:", command.name, err)
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

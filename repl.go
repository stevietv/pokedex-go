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

		command := input[0]

		switch command {
		case "exit":
			commandExit()
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	return strings.Fields(text)
}

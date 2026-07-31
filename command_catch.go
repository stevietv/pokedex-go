package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You need to provide a pokemon name")
	}

	name := args[0]

	pokemonResponse, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonResponse.Name)

	if !catchSucessful(pokemonResponse.BaseExperience) {
		fmt.Printf("%v escaped!\n", pokemonResponse.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemonResponse.Name)
	cfg.pokedex[pokemonResponse.Name] = pokemonResponse

	return nil
}

func catchSucessful(baseExperience int) bool {
	chanceToCatch := rand.IntN(1000)
	return chanceToCatch > baseExperience
}

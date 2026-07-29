package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	mapResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = mapResponse.Next
	cfg.prevLocations = mapResponse.Previous

	for _, location := range mapResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.prevLocations == nil {
		return errors.New("you're on the first page")
	}

	mapResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = mapResponse.Next
	cfg.prevLocations = mapResponse.Previous

	for _, location := range mapResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

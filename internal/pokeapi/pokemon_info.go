package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (PokemonResponse, error) {
	if name == "" {
		return PokemonResponse{}, errors.New("Please provide a valid pokemon name")
	}

	url := baseUrl + "/pokemon/" + name

	body, found := c.cache.Get(url)

	if !found {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonResponse{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return PokemonResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		if err != nil {
			return PokemonResponse{}, err
		}
	}

	pokemonResponse := PokemonResponse{}
	err := json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	if !found {
		c.cache.Add(url, body)
	}

	return pokemonResponse, nil

}

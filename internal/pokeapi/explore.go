package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (LocationResponse, error) {
	url := baseUrl + "/location-area/" + location

	body, found := c.cache.Get(url)

	if !found {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationResponse{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationResponse{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return LocationResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		if err != nil {
			return LocationResponse{}, err
		}
	}

	locationResponse := LocationResponse{}
	err := json.Unmarshal(body, &locationResponse)
	if err != nil {
		return LocationResponse{}, err
	}

	if !found {
		c.cache.Add(url, body)
	}

	return locationResponse, nil

}

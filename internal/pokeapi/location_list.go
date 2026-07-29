package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (MapResponse, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MapResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return MapResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return MapResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return MapResponse{}, err
	}

	mapResponse := MapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return MapResponse{}, err
	}

	return mapResponse, nil

}

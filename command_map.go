package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {

	if nextLocations == "" {
		return errors.New("no more locations")
	}

	res, err := http.Get(nextLocations)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return err
	}

	mapResponse := mapResponse{}
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return err
	}

	if mapResponse.Next != nil {
		nextLocations = *mapResponse.Next
	} else {
		nextLocations = ""
	}

	if mapResponse.Previous != nil {
		prevLocations = *mapResponse.Previous
	} else {
		prevLocations = ""
	}

	for _, location := range mapResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

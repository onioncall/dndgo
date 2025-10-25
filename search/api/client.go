package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url = "https://www.dnd5eapi.co/api/2014"

type PathType string

type BaseRequest struct {
	Name     string
	PathType PathType
}

func ExecuteGetRequest[T any](p PathType, criteria string) (T, error) {
	path := fmt.Sprintf("%s/%s/%s", url, p, criteria)
	resp, err := http.Get(path)
	if err != nil {
		return *new(T), fmt.Errorf("Failed Request: %s, Error: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return *new(T), fmt.Errorf("Failed Request: %s, HTTP Status: %d", path, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(T), fmt.Errorf("Request: %s, Read Response Failed: %w", path, err)
	}

	var obj T
	err = json.Unmarshal([]byte(body), &obj)
	if err != nil {
		return *new(T), fmt.Errorf("Request: %s, Unmarshal Response Failed: %w", path, err)
	}

	return obj, nil
}

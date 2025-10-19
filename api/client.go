package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/onioncall/dndgo/logger"
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
		logErr := fmt.Errorf("Failed Request: %s, Error: %v", path, err)
		logger.HandleError(err, logErr)

		return *new(T), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Failed Request: %s, HTTP Status: %v", path, resp.StatusCode)
		return *new(T), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Request: %s, Read Response Failed: %v", path, err)
		return *new(T), err
	}

	var obj T
	err = json.Unmarshal([]byte(body), &obj)
	if err != nil {
		err := fmt.Errorf("Request: %s, Unmarshal Response Failed: %v", path, err)
		return *new(T), err
	}

	return obj, nil
}

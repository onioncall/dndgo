package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

var url = "https://www.dnd5eapi.co/api/2014"

type PathType string

type BaseRequest struct {
	Name     string
    PathType PathType
}

func ExecuteGetRequest[T any](p PathType, criteria string) (T, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", url, p, criteria))
	if err != nil {
		fmt.Println("Request Failed")
		panic(err)
	}
	defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Response Failed")
		return *new(T), err
	}

	var obj T
    err = json.Unmarshal([]byte(body), &obj)
	if err != nil {
		fmt.Println("Unmarshal Response Failed")
		return *new(T), err
	}

	return obj, nil
}

func RequestFactory[T any](args string, request T, p PathType) T {
    v := reflect.ValueOf(&request).Elem()
    
    v.FieldByName("PathType").Set(reflect.ValueOf(p))
	v.FieldByName("Name").SetString(args)
    
    return request
}

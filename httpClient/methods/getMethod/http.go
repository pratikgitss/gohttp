package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resData []User

	err = json.Unmarshal(data, &resData)

	if err != nil {
		return nil, err
	}

	return resData, nil

}

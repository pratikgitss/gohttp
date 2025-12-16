package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, nil
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}

	res, err := client.Do(req)

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return User{}, nil
	}

	var resData User

	err = json.Unmarshal(body, &resData)

	if err != nil {
		return User{}, nil
	}

	return resData, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return User{}, nil
	}

	req.Header.Set("X-API-key", apiKey)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return User{}, nil
	}
	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	var resData User

	err = json.Unmarshal(body, &resData)

	if err != nil {
		return User{}, nil
	}

	return resData, nil

}

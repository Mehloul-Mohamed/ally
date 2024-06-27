package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetChallList(url string, apiToken string) ([]byte, error) {
	r, err := http.NewRequest("GET", url+"/api/v1/challenges", nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{Timeout: time.Second * 10}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", apiToken)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}

func GetChallenge(id int, url string, apiToken string) (*CtfdChallResponse, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/challenges/%d", url, id), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: time.Second * 10}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", apiToken)

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var challenge CtfdChallResponse
	err = json.Unmarshal(responseBytes, &challenge)
	if err != nil {
		return nil, err
	}
	return &challenge, nil
}

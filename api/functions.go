package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetChallList(url string, apiToken string) (*CtfdChallListResponse, error) {
	// If no url or token are given read them from the credentials file
	if url == "" && apiToken == "" {
		wd, _ := os.Getwd()
		bytes, err := os.ReadFile(wd + "/credentials.txt")
		if err != nil {
			return nil, err
		}
		slice := strings.Split(string(bytes), "\n")
		url = slice[0]
		apiToken = slice[1]
	}

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

	// Parse Json
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var challs CtfdChallListResponse
	err = json.Unmarshal(responseBytes, &challs)
	if err != nil {
		return nil, err
	}
	return &challs, nil
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

	// Parse Json
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

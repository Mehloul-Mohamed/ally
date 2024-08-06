package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetChallList(url string, apiToken string) (*CtfdChallListResponse, error) {
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

func GetTeamInfo(url string, apiToken string) (*CtfdTeamResponse, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/teams/me", url), nil)
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
	var team CtfdTeamResponse
	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func GetTopThree(url string, apiToken string) (*CtfdScoreBoardResponse, error) {
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/scoreboard/top/3", url), nil)
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
	var scoreBoard CtfdScoreBoardResponse
	err = json.Unmarshal(responseBytes, &scoreBoard)
	if err != nil {
		return nil, err
	}
	return &scoreBoard, nil
}

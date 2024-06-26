package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetChallList(url string, apiToken string) ([]byte, error) {
	r, err := http.NewRequest("Get", url+"/api/v1/challenges", nil)
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
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	f, err := os.Create(wd + "/challenges.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	f.Write(responseBytes)
	return responseBytes, nil
}

// Just realized with the default config I don't have permission to get the challenges files from
// the challenge files endpoint as a regular user.... ahhhhhhhh
// I found that the challenge endpoint gave me what I need, but this functionality seems undocumented?
// For now this is WorksOnMyMachine™ certified
func GetChallengeFiles(id int, url string, apiToken string) (*[]string, error) {
	r, err := http.NewRequest("Get", fmt.Sprintf("%s/api/v1/challenges/%d", url, id), nil)
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
	return &challenge.Data.Files, nil
}

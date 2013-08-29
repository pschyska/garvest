package lib

import (
	"net/http"
	"io/ioutil"
)

type Harvest struct{}

func (h Harvest) Connect() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pschyska.harvestapp.com/daily", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("foo@bar.com", "password123")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "garvest")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

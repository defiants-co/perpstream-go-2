package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Fetch[T any](path string, apiUrl string, params map[string]string, result *T) error {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return errors.New("failed to parse URL")
	}
	u.Path = path
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return errors.New("failed to fetch")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return errors.New(fmt.Sprintf("received status code %d: %s", resp.StatusCode, bodyString))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading response body: %v", err))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return errors.New(fmt.Sprintf("error unmarshalling response body: %v", err))
	}

	return nil
}

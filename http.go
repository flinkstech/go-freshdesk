package freshdesk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *apiClient) postJSON(path string, requestBody interface{}, out interface{}) error {
	httpClient := &http.Client{}
	jsonb, _ := json.Marshal(&requestBody)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s.freshdesk.com%s", c.domain, path), bytes.NewReader(jsonb))

	req.SetBasicAuth(c.apiKey, "X")
	req.Header.Add("Content-type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("Freshdesk server didn't like the request")
	}

	err = json.NewDecoder(res.Body).Decode(out)

	return err
}

func (c *apiClient) get(path string, out interface{}) (http.Header, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.freshdesk.com%s", c.domain, path), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "X")
	req.Header.Add("Content-type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("received status code %d (200 expected)", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(out)

	return res.Header, err
}

func (c *apiClient) getNextLink(headers http.Header) (string, bool) {
	link := headers.Get("link")
	if link != "" {
		return strings.TrimPrefix(strings.TrimSuffix(link, ">; rel=\"next\""), fmt.Sprintf("<https://%s.freshdesk.com", c.domain)), true
	}
	return "", false
}

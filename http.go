package freshdesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *apiClient) postJSON(path string, requestBody []byte, out interface{}, expectedStatus int) error {
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s.freshdesk.com%s", c.domain, path), bytes.NewReader(requestBody))

	req.SetBasicAuth(c.apiKey, "X")
	req.Header.Add("Content-type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != expectedStatus {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		return fmt.Errorf("received status code %d (200 expected)", res.StatusCode)
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

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
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			body, _ := ioutil.ReadAll(res.Body)
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, body, "", "\t")
			if c.logger != nil {
				c.logger.Println(string(jsonBuffer.Bytes()))
			}
		}
		return fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus)
	}

	err = json.NewDecoder(res.Body).Decode(out)

	return err
}

func (c *apiClient) put(path string, requestBody []byte, out interface{}, expectedStatus int) error {
	httpClient := &http.Client{}
	if c.logger != nil {
		c.logger.Println(string(requestBody))
	}
	req, _ := http.NewRequest("PUT", fmt.Sprintf("https://%s.freshdesk.com%s", c.domain, path), bytes.NewReader(requestBody))

	req.SetBasicAuth(c.apiKey, "X")
	req.Header.Add("Content-type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != expectedStatus {
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			body, _ := ioutil.ReadAll(res.Body)
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, body, "", "\t")
			if c.logger != nil {
				c.logger.Println(string(jsonBuffer.Bytes()))
			}
		}
		return fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus)
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

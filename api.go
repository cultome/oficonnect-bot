package oficonnectbot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct{}

func BuildClient() *Client {
	return &Client{}
}

func (c *Client) Get(url string, response any) error {
	return request(url, "GET", nil, response)
}

func (c *Client) Post(url string, payload io.Reader, response any) error {
	return request(url, "POST", payload, response)
}

func request(url, method string, payload io.Reader, response any) error {
	client := &http.Client{}

	req, err := http.NewRequest(string(method), url, payload)

	if err != nil {
		return fmt.Errorf("[-] Unable to create request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OFICONNECT_TOKEN")))

	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("[-] Problems with request: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, response)

	if err != nil {
		return fmt.Errorf("[-] Unable to parse response body: %s", err.Error())
	}

	return nil
}

package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

type Result struct {
	Items []struct {
		ID              int      `json:"id"`
		Name            string   `json:"name"`
		FullName        string   `json:"full_name"`
		URL             string   `json:"url"`
		HTMLURL         string   `json:"html_url"`
		CloneURL        string   `json:"clone_url"`
		Description     string   `json:"description"`
		StargazersCount int      `json:"stargazers_count"`
		Watchers        int      `json:"watchers"`
		Topics          []string `json:"topics"`
		Language        string   `json:"language"`
		CreatedAt       string   `json:"created_at"`
		UpdatedAt       string   `json:"updated_at"`
	} `json:"items"`
}

func (result *Result) Draw(writer io.Writer) error {
	for _, item := range result.Items {
		fmt.Fprintf(writer, " %s\n", item.FullName)
	}
	return nil
}

func NewClient(rawBaseURL string) (*Client, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
	}, nil
}

func (client *Client) GetTopicItemList() (*Result, error) {
	req, err := http.NewRequest("GET", client.BaseURL.String()+"?q=topic", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.mercy-preview+json")
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result *Result
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

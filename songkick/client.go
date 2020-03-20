package songkick

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl string
	ApiKey string

	httpClient *http.Client
}

func NewClient(baseurl string, apiKey string) *Client {
	return &Client{
		BaseUrl: baseurl,
		ApiKey: apiKey,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) GetEvent(id string) (*Event, error) {

	url, err := url.Parse(c.BaseUrl + id + ".json")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", c.ApiKey)

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	e := &SongkickEventResponse{}

	if err = json.NewDecoder(resp.Body).Decode(e); err != nil {
		return nil, err
	}

	if e.ResultsPage.Status != "ok" {
		return nil, fmt.Errorf("failed to get event: %w", errors.New(e.ResultsPage.Error.Message))
	}

	return &e.ResultsPage.Results.Event, err

}
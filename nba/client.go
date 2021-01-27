package nba

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"fmt"
)

const baseURL = "api-nba-v1.p.rapidapi.com"

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Key        string
	Logger     *log.Logger
}

func New(Key string, logger *log.Logger) (*Client, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Key:        Key,
		Logger:     logger,
	}, nil
}

func (c *Client) NewRequest(ctx context.Context, method, relativePath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	url := "https://api-nba-v1.p.rapidapi.com/seasons/"

	req, _ := http.NewRequest("GET", url, nil)

    // set header
	req.Header.Add("x-rapidapi-host", baseURL)
	req.Header.Add("x-rapidapi-key", c.Key)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}


	
	
	// reqURL := *c.BaseURL

	// set path
	// reqURL.Path = path.Join(reqURL.Path, relativePath)
	// url := "https://api-nba-v1.p.rapidapi.com/seasons/"
	// reqURL.Path = url
	// set query
	// if queries != nil {
	// 	q := reqURL.Query()
	// 	for k, v := range queries {
	// 		q.Add(k, v)
	// 	}
	// 	reqURL.RawQuery = q.Encode()
	// }

	// instantiate request
	// req, err := http.NewRequest(method, reqURL.String(), reqBody)
	// if err != nil {
	// 	return nil, err
	// }

	

	// set context
	// req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) DoRequest(req *http.Request, respBody interface{}) (int, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return resp.StatusCode, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(bodyBytes, respBody); err != nil {
		fmt.Printf("%#v\n", err)
		return 0, err
	}

	return resp.StatusCode, nil
}


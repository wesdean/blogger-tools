package blogger_tools_blogger

import (
	"errors"
	"fmt"
	"github.com/google/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var bloggerBaseURL = "https://www.googleapis.com/blogger/v3/blogs/"

type Client struct {
	logger     *logger.Logger
	baseURL    string
	apiKey     string
	blogId     string
	httpClient http.Client
}

func NewClient(logger *logger.Logger, apiKey string, blogId string) *Client {
	return &Client{
		logger:     logger,
		baseURL:    bloggerBaseURL + blogId,
		apiKey:     apiKey,
		blogId:     blogId,
		httpClient: http.Client{},
	}
}

func (client *Client) SendRequest(path string) ([]byte, error) {
	options := url.Values{}
	options.Set("key", client.apiKey)
	resp, err := client.httpClient.Get(client.baseURL + path + "?" + options.Encode())
	if err != nil {
		client.logger.Error(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wrong response status: %v", resp.StatusCode))
	}

	closeBody := func(client *Client, body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			client.logger.Error(err)
		}
	}

	defer closeBody(client, resp.Body)
	return ioutil.ReadAll(resp.Body)
}

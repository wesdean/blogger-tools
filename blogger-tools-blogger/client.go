package blogger_tools_blogger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var bloggerBaseURL = "https://www.googleapis.com/blogger/v3/blogs/"

type Client struct {
	logger      *logger.Logger
	baseURL     string
	accessToken string
	blogId      string
	httpClient  *http.Client
}

type ErrorResponse struct {
	Error *ErrorResponseError `json:"error,omitempty"`
}

type ErrorResponseError struct {
	Errors  []ErrorResponseErrorItem `json:"errors,omitempty"`
	Code    int                      `json:"code,omitempty"`
	Message string                   `json:"message,omitempty"`
}

type ErrorResponseErrorItem struct {
	Domain  string `json:"domain,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewClient(logger *logger.Logger, accessToken string, blogId string) *Client {
	return &Client{
		logger:      logger,
		baseURL:     bloggerBaseURL + blogId,
		accessToken: accessToken,
		blogId:      blogId,
		httpClient:  &http.Client{},
	}
}

func (client *Client) SendRequest(path string, params map[string]string) ([]byte, error) {
	options := url.Values{}
	if params != nil {
		for key, value := range params {
			options.Set(key, value)
		}
	}
	if client.accessToken != "" {
		options.Set("access_token", client.accessToken)
	} else {
		options.Set("access_token", os.Getenv("ACCESS_TOKEN"))
	}

	urlStr := client.baseURL + path + "?" + options.Encode()
	resp, err := client.httpClient.Get(urlStr)
	if err != nil {
		client.logger.Error(err)
		return nil, err
	}

	closeBody := func(client *Client, body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			client.logger.Error(err)
		}
	}

	defer closeBody(client, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		respError := ErrorResponse{}
		err = json.Unmarshal(body, &respError)
		if err != nil {
			client.logger.Error(err)
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("%v: %v", resp.StatusCode, respError.Error.Message))
	}

	return body, nil
}

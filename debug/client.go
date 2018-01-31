// Package debug just contains some junk to help me debug the API while in
// development. Not really meaningful anywhere else.
package debug

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Client struct {
	*http.Client
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, nil
	}

	fmt.Println(string(b))

	return resp, err
}

type MockClient struct{}

func (c *MockClient) Do(r *http.Request) (*http.Response, error) {
	b, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		return nil, nil
	}

	fmt.Println(string(b))

	return nil, errors.New("mock client")
}

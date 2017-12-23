package hue

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const base = "http://%s/api/%s"

// Client describes an interface that consumes an HTTP request to produce a
// response.
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// Hue is a client to a Phillips Hue bridge. It allows control of Hue lights
// using the RESTful API.
type Hue struct {
	client Client
	ip     string
	user   string
}

// New returns a Hue client. It connects to the bridge at the provided IP, using
// the given user ID, using the provided Client.
func New(ip, user string, client Client) *Hue {
	return &Hue{ip: ip, user: user, client: client}
}

// request is just a wrapper around the http package that does all the required
// setup.
func (h *Hue) request(method, frag string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf(base, h.ip, h.user) + frag

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "could not build request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	return resp, err
}

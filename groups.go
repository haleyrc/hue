package hue

import (
	"encoding/json"
	"io"
	"net/http/httputil"
	"time"

	"github.com/pkg/errors"
)

// Groups returns a list of user-defined light groups.
func (h *Hue) Groups() ([]*Group, error) {
	resp, err := h.request("GET", "/groups", nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not get groups")
	}

	var body map[string]*Group
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "could not decode groups response")
	}
	resp.Body.Close()

	groups := make([]*Group, 0)
	for id, group := range body {
		group.ID = id
		groups = append(groups, group)
	}

	return groups, nil
}

// SetAll applies the provided StateMod to the master group that contains all the
// lights in the system.
func (h *Hue) SetAll(state StateMod) error {
	pr, pw := io.Pipe()
	defer pr.Close()

	errs := make(chan error, 1)
	go func() {
		errs <- json.NewEncoder(pw).Encode(state)
		pw.Close()
		close(errs)
	}()

	resp, err := h.request("PUT", "/groups/0/action", pr)
	if err != nil {
		return errors.Wrap(err, "could not set state")
	}
	httputil.DumpResponse(resp, true)

	for {
		select {
		case err := <-errs:
			return errors.Wrap(err, "could not set state")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

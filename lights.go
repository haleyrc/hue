package hue

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Light returns the information and state for the Phillips Hue bulb with the
// given ID.
func (h *Hue) Light(id string) (*Light, error) {
	frag := fmt.Sprintf("/lights/%s", id)
	resp, err := h.request("GET", frag, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not get light")
	}

	var body *Light
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "could not decode light response")
	}
	resp.Body.Close()

	body.ID = id

	return body, nil
}

// Lights returns a list of all reachable Phillips Hue lightbulbs and their
// states.
func (h *Hue) Lights() ([]*Light, error) {
	resp, err := h.request("GET", "/lights", nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not get lights")
	}

	var body map[string]*Light
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "could not decode lights response")
	}
	resp.Body.Close()

	lights := make([]*Light, 0)
	for id, light := range body {
		if !light.State.Reachable {
			continue
		}
		light.ID = id
		lights = append(lights, light)
	}

	return lights, nil
}

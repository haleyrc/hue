package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
)

// TODO (RCH): Add transition time modifier

func On() StateMod {
	return StateMod{"on": true}
}

func Off() StateMod {
	return StateMod{"on": false}
}

func NewState(opts ...StateOption) StateMod {
	var m = make(StateMod)

	for _, opt := range opts {
		opt(m)
	}

	return m
}

type StateMod map[string]interface{}
type StateOption func(StateMod)

func WithBrightness(bri int) StateOption {
	return func(m StateMod) {
		if bri < 1 {
			bri = 1
		}

		if bri > 254 {
			bri = 254
		}

		m["bri"] = bri
	}
}

func WithHue(hue int) StateOption {
	return func(m StateMod) {
		if hue < 0 {
			hue = 0
		}

		if hue > 65535 {
			hue = 65535
		}

		m["hue"] = hue
	}
}

func WithSaturation(sat int) StateOption {
	return func(m StateMod) {
		if sat < 0 {
			sat = 0
		}

		if sat > 254 {
			sat = 254
		}

		m["sat"] = sat
	}
}

type Alert int

const (
	CancelAlert = iota
	ShortAlert
	LongAlert
)

func WithAlert(alert Alert) StateOption {
	return func(m StateMod) {
		switch alert {
		case ShortAlert:
			m["alert"] = "select"
		case LongAlert:
			m["alert"] = "lselect"
		case CancelAlert:
			m["alert"] = "none"
		default:
			m["alert"] = "none"
		}
	}
}

type Effect int

const (
	CancelEffect = iota
	ColorLoop
)

func WithEffect(effect Effect) StateOption {
	return func(m StateMod) {
		switch effect {
		case ColorLoop:
			m["effect"] = "colorloop"
		case CancelEffect:
			m["effect"] = "none"
		default:
			m["effect"] = "none"
		}
	}
}

func (h *Hue) SetState(id string, s StateMod) error {
	pr, pw := io.Pipe()
	defer pr.Close()

	errs := make(chan error, 1)
	go func() {
		errs <- json.NewEncoder(pw).Encode(s)
		pw.Close()
		close(errs)
	}()

	_, err := h.request("PUT", fmt.Sprintf("/lights/%s/state", id), pr)
	if err != nil {
		return errors.Wrap(err, "could not set state")
	}

	for {
		select {
		case err := <-errs:
			return errors.Wrap(err, "could not set state")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

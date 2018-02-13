package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
)

// On returns a StateMod setting the "on" state to true.
func On() StateMod {
	return StateMod{"on": true}
}

// Off returns a StateMod setting the "on" state to false.
func Off() StateMod {
	return StateMod{"on": false}
}

// NewState returns a StateMod that is the combination of the provided options.
// Options are processed in order, so passing two of the same option ends up in
// a last write wins scenario.
func NewState(opts ...StateOption) StateMod {
	var m = make(StateMod)

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// StateMod is a structure suitable for converting to JSON that can be passed to
// the state update endpoint for a single light. This can be created manually,
// but it's ideal to use the StateOption functions since they do some sanity
// checks beforehand.
type StateMod map[string]interface{}

// StateOption is a function that modifies the given StateMod. Most settings are
// specified as "generator" function that return a StateOption function when
// passed a value for whatever parameter is being modified.
type StateOption func(StateMod)

// WithBrightness returns a StateOption function that sets the "bri" parameter
// to the given value. The value is clamped to be between 1 and 254, with 1
// being the lowest a bulb can operate, and 254 being maximum brightness.
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

// WithHue returns a StateOption function that sets the "hue" parameter to the
// given vavlue. The value is clamped to be between 0 and 65535, with both being
// red, and additional hues occupying the space between.
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

// WithSaturation returns a StateOption function that sets the "sat" parameter
// to the given value. The value is clamped to be between 0 and 254, with 254
// being maximum color saturation, and 0 being white.
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

// WithTransitionTime returns a StateOption function that sets the
// "transitiontime" parameter to the given value. The value is clamped to be
// greater than 0, since negative transition times are non-sensical.
func WithTransitionTime(t int) StateOption {
	return func(m StateMod) {
		if t < 0 {
			t = 0
		}

		m["transitiontime"] = t
	}
}

// Alert represents one of the recognized alert states.
type Alert int

const (
	// CancelAlert represents an alert status of "none".
	CancelAlert = iota
	// ShortAlert represents an alert status of "select", which is a single flash.
	ShortAlert
	// LongAlert represents an alert status of "lselect", which pulses for 15s or
	// until "none" is set.
	LongAlert
)

// WithAlert returns a StateOption that sets the "alert" parameter to the value
// corresponding to the provided Alert. Unrecognized Alerts are assumed to be
// "none".
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

// Effect represents one of the recognized effect states.
type Effect int

const (
	// CancelEffect represents an effect status of "none".
	CancelEffect = iota
	// ColorLoop represents an effect status of "colorloop".
	ColorLoop
)

// WithEffect returns a StateOption that sets the "effect" parameter to the
// value corresponding to the provided Effect. Unrecognized Effects are assumed
// to be "none".
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

// SetState sets the active state of the provided light to the values supplied
// in the given StateMod.
func (h *Hue) SetState(l *Light, s StateMod) error {
	pr, pw := io.Pipe()
	defer pr.Close()

	errs := make(chan error, 1)
	go func() {
		errs <- json.NewEncoder(pw).Encode(s)
		pw.Close()
		close(errs)
	}()

	if _, err := h.request("PUT", fmt.Sprintf("/lights/%s/state", l.ID), pr); err != nil {
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

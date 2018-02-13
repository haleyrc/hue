package hue_test

import (
	"testing"

	"github.com/haleyrc/hue"
)

func TestWithBrightness(t *testing.T) {
	var m = make(hue.StateMod)

	if bri, ok := m["bri"]; ok {
		t.Errorf("expected brightness not to be set, but it was %d", bri)
	}

	tests := []struct{ In, Want int }{
		{In: 10, Want: 10},
		{In: 0, Want: 1},
		{In: -1, Want: 1},
		{In: 255, Want: 254},
		{In: 1000, Want: 254},
	}

	for _, test := range tests {
		hue.WithBrightness(test.In)(m)

		bri, ok := m["bri"]
		if !ok {
			t.Errorf("expected brightness to be set, but it wasn't")
		}

		if bri != test.Want {
			t.Errorf("expected brightness to be %d, but it was %d", test.Want, bri)
		}
	}
}

func TestWithHue(t *testing.T) {
	var m = make(hue.StateMod)

	if hue, ok := m["hue"]; ok {
		t.Errorf("expected hue not to be set, but it was %d", hue)
	}

	tests := []struct{ In, Want int }{
		{In: 0, Want: 0},
		{In: 65535, Want: 65535},
		{In: 20000, Want: 20000},
		{In: -1, Want: 0},
		{In: 65536, Want: 65535},
	}

	for _, test := range tests {
		hue.WithHue(test.In)(m)

		hue, ok := m["hue"]
		if !ok {
			t.Errorf("expected hue to be set, but it wasn't")
		}

		if hue != test.Want {
			t.Errorf("expected hue to be %d, but it was %d", test.Want, hue)
		}
	}
}

func TestWithSaturation(t *testing.T) {
	var m = make(hue.StateMod)

	if sat, ok := m["sat"]; ok {
		t.Errorf("expected saturation not to be set, but it was %d", sat)
	}

	tests := []struct{ In, Want int }{
		{In: 0, Want: 0},
		{In: 254, Want: 254},
		{In: 100, Want: 100},
		{In: -1, Want: 0},
		{In: 255, Want: 254},
	}

	for _, test := range tests {
		hue.WithSaturation(test.In)(m)

		sat, ok := m["sat"]
		if !ok {
			t.Errorf("expected saturation to be set, but it wasn't")
		}

		if sat != test.Want {
			t.Errorf("expected saturation to be %d, but it was %d", test.Want, sat)
		}
	}
}

func TestWithTransitionTime(t *testing.T) {
	var m = make(hue.StateMod)

	if tt, ok := m["transitiontime"]; ok {
		t.Errorf("expected transition time not to be set, but it was %d", tt)
	}

	tests := []struct{ In, Want int }{
		{In: 0, Want: 0},
		{In: 30, Want: 30},
		{In: -1, Want: 0},
	}

	for _, test := range tests {
		hue.WithTransitionTime(test.In)(m)

		tt, ok := m["transitiontime"]
		if !ok {
			t.Errorf("expected transition time to be set, but it wasn't")
		}

		if tt != test.Want {
			t.Errorf("expected transition time to be %d, but it was %d", test.Want, tt)
		}
	}
}

func TestWithAlert(t *testing.T) {
	var m = make(hue.StateMod)

	if alert, ok := m["alert"]; ok {
		t.Errorf("expected alert not to be set, but it was %s", alert)
	}

	tests := []struct {
		In   hue.Alert
		Want string
	}{
		{In: hue.ShortAlert, Want: "select"},
		{In: hue.LongAlert, Want: "lselect"},
		{In: hue.CancelAlert, Want: "none"},
		{In: hue.Alert(-1), Want: "none"},
	}

	for _, test := range tests {
		hue.WithAlert(test.In)(m)

		alert, ok := m["alert"]
		if !ok {
			t.Errorf("expected alert to be set, but it wasn't")
		}

		if alert != test.Want {
			t.Errorf("expected alert to be %s, but it was %s", test.Want, alert)
		}
	}
}

func TestWithEffect(t *testing.T) {
	var m = make(hue.StateMod)

	if effect, ok := m["effect"]; ok {
		t.Errorf("expected effect not to be set, but it was %s", effect)
	}

	tests := []struct {
		In   hue.Effect
		Want string
	}{
		{In: hue.ColorLoop, Want: "colorloop"},
		{In: hue.CancelEffect, Want: "none"},
		{In: hue.Effect(-1), Want: "none"},
	}

	for _, test := range tests {
		hue.WithEffect(test.In)(m)

		effect, ok := m["effect"]
		if !ok {
			t.Errorf("expected effect to be set, but it wasn't")
		}

		if effect != test.Want {
			t.Errorf("expected effect to be %s, but it was %s", test.Want, effect)
		}
	}
}

type NewStateTest struct {
	Opts []hue.StateOption
	Want hue.StateMod
}

func TestNewState(t *testing.T) {
	tests := []NewStateTest{
		NewStateTest{
			Opts: []hue.StateOption{hue.WithBrightness(10)},
			Want: hue.StateMod{"bri": 10},
		},
		NewStateTest{
			Opts: []hue.StateOption{hue.WithHue(10)},
			Want: hue.StateMod{"hue": 10},
		},
		NewStateTest{
			Opts: []hue.StateOption{hue.WithSaturation(10)},
			Want: hue.StateMod{"sat": 10},
		},
		NewStateTest{
			Opts: []hue.StateOption{hue.WithAlert(hue.ShortAlert)},
			Want: hue.StateMod{"alert": "select"},
		},
		NewStateTest{
			Opts: []hue.StateOption{hue.WithEffect(hue.ColorLoop)},
			Want: hue.StateMod{"effect": "colorloop"},
		},
		NewStateTest{
			Opts: []hue.StateOption{hue.WithTransitionTime(30)},
			Want: hue.StateMod{"transitiontime": 30},
		},
		NewStateTest{
			Opts: []hue.StateOption{
				hue.WithBrightness(10),
				hue.WithEffect(hue.ColorLoop),
			},
			Want: hue.StateMod{"bri": 10, "effect": "colorloop"},
		},
		NewStateTest{
			Opts: []hue.StateOption{
				hue.WithHue(10),
				hue.WithSaturation(20),
			},
			Want: hue.StateMod{"hue": 10, "sat": 20},
		},
		NewStateTest{
			Opts: []hue.StateOption{
				hue.WithBrightness(10),
				hue.WithHue(20),
				hue.WithSaturation(30),
				hue.WithAlert(hue.LongAlert),
				hue.WithEffect(hue.ColorLoop),
				hue.WithTransitionTime(30),
			},
			Want: hue.StateMod{
				"bri":            10,
				"hue":            20,
				"sat":            30,
				"alert":          "lselect",
				"effect":         "colorloop",
				"transitiontime": 30,
			},
		},
	}

	for _, test := range tests {
		m := hue.NewState(test.Opts...)
		for k, want := range test.Want {
			got, ok := m[k]
			if !ok {
				t.Errorf("%s missing from state", k)
				continue
			}

			if want != got {
				t.Errorf("%s mismatch. wanted %+v, got %+v", k, want, got)
			}
		}

		for k := range m {
			if _, ok := test.Want[k]; !ok {
				t.Errorf("extra option %s in state", k)
			}
		}
	}
}

func TestOn(t *testing.T) {
	sm := hue.On()
	if !sm["on"].(bool) {
		t.Errorf(`expected "on" to be true, but it wasn't`)
	}
}

func TestOff(t *testing.T) {
	sm := hue.Off()
	if sm["on"].(bool) {
		t.Errorf(`expected "off" to be true, but it wasn't`)
	}
}

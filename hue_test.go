package hue_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/haleyrc/hue"
)

func TestLights(t *testing.T) {
	ip := os.Getenv("HUE_BRIDGE_IP")
	id := os.Getenv("HUE_USER_ID")
	hc := hue.New(ip, id, &http.Client{})

	ls, err := hc.Lights()
	if err != nil {
		t.Errorf("error getting lights state: %v\n", err)
		t.FailNow()
	}

	for _, l := range ls {
		t.Logf("%v\n", l)
	}
}

func TestGroups(t *testing.T) {
	ip := os.Getenv("HUE_BRIDGE_IP")
	id := os.Getenv("HUE_USER_ID")
	hc := hue.New(ip, id, &http.Client{})

	gs, err := hc.Groups()
	if err != nil {
		t.Errorf("error getting lights state: %v\n", err)
		t.FailNow()
	}

	for _, g := range gs {
		t.Logf("%s\n", g)
	}

	if err := hc.SetAll(hue.NewState(hue.WithBrightness(1))); err != nil {
		t.Errorf("error setting state for all: %v\n", err)
	}
}

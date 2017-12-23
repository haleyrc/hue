package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/haleyrc/hue"
)

func main() {
	bridgeIP := os.Getenv("HUE_BRIDGE_IP")
	userID := os.Getenv("HUE_USER_ID")
	fmt.Println(userID)

	if bridgeIP == "" {
		log.Fatalln("bridge ip is required")
	}

	if userID == "" {
		log.Fatalln("user id is required")
	}

	h := hue.New(bridgeIP, userID, &DebugClient{&http.Client{}})

	lights, err := h.Lights()
	if err != nil {
		log.Fatalln(err)
	}

	for _, light := range lights {
		fmt.Println(light)
	}

	l, err := h.Light("1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(l)

	// l.On()
	// h.On(l)

	// l.Off()
	// h.Off(l)

	// l.SetBrightness(254)
	// h.SetBrightness(l, 254)
	// h.SetAllBrightness(254)

	// l.SetHue(65535)
	// h.SetHue(l, 65535)

	// l.SetSaturation(254)
	// h.SetSaturation(l, 254)

	// l.Alert()
	// h.Alert(l)

	// l.Breathe()
	// h.Breathe(l)

	// l.SetTransitionTime(4)
	// h.SetTransitionTime(l, 4)

	// h.Commit([dryrun])
	// h.Explain()
}

type DebugClient struct {
	*http.Client
}

func (c *DebugClient) Do(r *http.Request) (*http.Response, error) {
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

// This program is a test of the API that also happens to be as festive as I
// get. Cycles the light with ID "1" between and green in a pleasant fashion.
// For me, this is an exterior light near my front door, making it okay that I
// didn't put up any other lights.
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/haleyrc/hue"
	"github.com/haleyrc/hue/debug"
)

func main() {
	bridgeIP := os.Getenv("HUE_BRIDGE_IP")
	userID := os.Getenv("HUE_USER_ID")

	if bridgeIP == "" {
		log.Fatalln("bridge ip is required")
	}

	if userID == "" {
		log.Fatalln("user id is required")
	}

	h := hue.New(bridgeIP, userID, &debug.Client{&http.Client{}})

	red := hue.NewState(
		hue.WithSaturation(255),
		hue.WithHue(0),
		hue.WithBrightness(128),
		hue.WithTransitionTime(30),
	)

	green := hue.NewState(
		hue.WithSaturation(255),
		hue.WithHue(25500),
		hue.WithBrightness(128),
		hue.WithTransitionTime(30),
	)

	l, err := h.Light("1")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		if err := h.SetState(l, red); err != nil {
			log.Fatalln(err)
		}
		<-time.After(4 * time.Second)

		if err := h.SetState(l, green); err != nil {
			log.Fatalln(err)
		}
		<-time.After(4 * time.Second)
	}
}

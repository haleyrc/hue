package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/haleyrc/hue"
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

	h := hue.New(bridgeIP, userID, &http.Client{})

	red := hue.NewState(
		hue.WithSaturation(255),
		hue.WithHue(0),
		hue.WithBrightness(255),
		hue.WithTransitionTime(30),
	)

	green := hue.NewState(
		hue.WithSaturation(255),
		hue.WithHue(25500),
		hue.WithBrightness(255),
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

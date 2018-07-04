package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/haleyrc/hue"
)

const (
	Orange = 5825
	Green  = 21663
)

func main() {
	var (
		bridgeIP = os.Getenv("HUE_BRIDGE_IP")
		userID   = os.Getenv("HUE_USER_ID")
	)

	if bridgeIP == "" {
		log.Fatalln("bridge ip is required")
	}

	if userID == "" {
		log.Fatalln("user id is required")
	}

	hc := hue.New(bridgeIP, userID, &http.Client{})

	currentState := 0
	states := []hue.StateMod{
		hue.NewState(hue.WithHue(Orange), hue.WithSaturation(255), hue.WithTransitionTime(20)),
		hue.NewState(hue.WithSaturation(0), hue.WithTransitionTime(20)),
		hue.NewState(hue.WithHue(Green), hue.WithSaturation(255), hue.WithTransitionTime(20)),
	}

	for {
		hc.SetAll(states[currentState])
		currentState = (currentState + 1) % len(states)
		time.Sleep(3 * time.Second)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haleyrc/hue"
)

var (
	Red    = hue.WithXY(hue.XY{0.675, 0.322})
	Orange = hue.WithXY(hue.XY{0.6173877280762861, 0.3644511477332629})
	Yellow = hue.WithHue(10699)
	Green  = hue.WithHue(24432)
	Blue   = hue.WithXY(hue.XY{0.167, 0.04})
	Violet = hue.WithHue(48089)
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

	hc := hue.New(bridgeIP, userID, nil)

	ls, err := hc.Lights()
	if err != nil {
		log.Fatalf("could not get lights: %v\n", err)
	}

	l, ok := ls.Get("Sitting Room Tall Lamp")
	if !ok {
		log.Fatalln("could not find light by name")
	}

	fmt.Printf("Running on %s...\n", l.Name)

	currentState := 0
	brightness := hue.WithBrightness(255)
	saturation := hue.WithSaturation(255)
	transition := hue.WithTransitionTime(1)
	states := []hue.StateMod{
		hue.NewState(Red, brightness, saturation, transition),
		hue.NewState(Orange, brightness, saturation, transition),
		hue.NewState(Yellow, brightness, saturation, transition),
		hue.NewState(Green, brightness, saturation, transition),
		hue.NewState(Blue, brightness, saturation, transition),
		hue.NewState(Violet, brightness, saturation, transition),
	}

	for {
		newState := states[currentState]
		if err := hc.SetState(l, newState); err != nil {
			log.Printf("error setting state: %v\n", err)
		}
		currentState = (currentState + 1) % len(states)
		time.Sleep(1000 * time.Millisecond)
	}
}

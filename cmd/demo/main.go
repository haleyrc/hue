// This package is a demo of the API in the hue package. Because ids are
// hard-coded to match a known light in my configuration, it is not guaranteed
// to work meaningfully anywhere else, but it should still serve as a decent
// demonstration of how to setup and use the API regardless.
package main

import (
	"fmt"
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

	h := hue.New(bridgeIP, userID, &debug.Client{Client: &http.Client{}})

	lights, err := h.Lights()
	if err != nil {
		log.Fatalln(err)
	}

	for _, light := range lights {
		fmt.Println(light)
	}

	l, err := h.Light("2")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(l)

	if err := h.SetState(l, hue.On()); err != nil {
		log.Fatalln(err)
	}

	<-time.After(2 * time.Second)
	redAlert := hue.NewState(
		hue.WithSaturation(255),
		hue.WithHue(0),
		hue.WithBrightness(255),
		hue.WithAlert(hue.LongAlert),
	)
	fmt.Printf("%+v\n", redAlert)
	if err := h.SetState(l, redAlert); err != nil {
		log.Fatalln(err)
	}

	<-time.After(5 * time.Second)
	resume := hue.NewState(hue.WithAlert(hue.CancelAlert), hue.WithBrightness(255), hue.WithSaturation(0))
	if err := h.SetState(l, resume); err != nil {
		log.Fatalln(err)
	}

	<-time.After(2 * time.Second)
	if err := h.SetState(l, hue.Off()); err != nil {
		log.Fatalln(err)
	}
}

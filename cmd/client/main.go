package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	l, err := h.Light("2")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(l)

	if err := h.SetState("2", hue.On()); err != nil {
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
	if err := h.SetState("2", redAlert); err != nil {
		log.Fatalln(err)
	}

	<-time.After(5 * time.Second)
	resume := hue.NewState(hue.WithAlert(hue.CancelAlert), hue.WithBrightness(255), hue.WithSaturation(0))
	if err := h.SetState("2", resume); err != nil {
		log.Fatalln(err)
	}

	<-time.After(2 * time.Second)
	if err := h.SetState("2", hue.Off()); err != nil {
		log.Fatalln(err)
	}
}

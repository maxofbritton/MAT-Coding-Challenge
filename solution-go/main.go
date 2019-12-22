/*
MAT Coding Challenge Solution
Written with Go

SPEED is solved by comparing the current positions and timestamps
POSITION is derived by ranking which driver has travelled the furthest (not very robust)

Overtakes are recognized when there is a change in position detected

Everything should be updated immediately, car vlues don't always arrive on order.
downside, this might cause artificial positions changes due to one car being updated beofre another

car previsou positions, distance covered, and speed reported every updated
position changes should be updated less frequently, but how often?

 */

package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type CarCoordinates struct {
	CarIndex int
	Location struct {
		Lat, Long float64
	}
	Timestamp uint64
}

type CarStatus struct {
	Timestamp, CarIndex, Type, Value string
}

type Events struct {
	Timestamp, CarIndex, Type, Value string
}

type PrevData struct {
	Timestamp uint64
	Location struct {
		Lat, Long float64
	}
	Distance float64
}

func process(car CarCoordinates, prev []PrevData) {

}

func main() {

	// Set initial
	const carCount = 6
	var prev [carCount]PrevData

	// How we should handle incoming messages
	var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

		var car CarCoordinates

		if err := json.Unmarshal([]byte(msg.Payload()), &car); err != nil {
			log.Fatal(err)
		}

		prev = process(car, prev)

		fmt.Printf("%+v\n", car)
		fmt.Printf("%f\n", car.Timestamp)
		fmt.Printf("%f\n", prev[car.CarIndex])


		// set
		if prev[coords.CarIndex].Distance == 0 {
			prev[coords.CarIndex].Location = coords.Location
			prev[coords.CarIndex].Timestamp = coords.Timestamp
		}

		// check if

		// get the timestamp and current location

		// get the previous timestanp and location

		// calculate delta time

		// Calculate distance

		// get total distance

		// calculate speed

		// now we want to update the previous location, and update total distance


		// text := fmt.Sprintf("this is result msg !")
		// token := client.Publish("carStatus", 0, false, text)
		// token.Wait()



	}






	// no clue what this does
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Create a new client?
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("StatusReport")
	// Set our function as the message handler
	opts.SetDefaultPublishHandler(f)

	// MQTT topic
	topic := "carCoordinates"

	// not sure how this works but it subscribes to the topic from before
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	// actually create the client and connect it
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}

	// wtf is this?
	<-c
}
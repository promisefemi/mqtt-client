package main

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"log"
	"mqtt-client/setup"
	"mqtt-client/subscribers"
)

func main() {
	client, err := setup.Setup()
	if err != nil {
		panic(err)
	}

	topicSubscribers := subscribers.RegisterSubscribers()
	for _, sub := range topicSubscribers {
		client.Router.RegisterHandler(sub.GetTopic(), sub.Handler)
		if _, err := client.Subscribe(context.Background(), &paho.Subscribe{
			Subscriptions: map[string]paho.SubscribeOptions{
				sub.GetTopic(): {QoS: sub.GetQoS()},
			},
		}); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Subscribed to %s \n", sub.GetTopic())
	}

	fmt.Println("app is running")
	<-make(chan bool)
}

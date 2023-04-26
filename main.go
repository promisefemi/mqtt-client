package main

import (
	"fmt"
	"mqtt-client/setup"
	"mqtt-client/subscribers"
)

func main() {
	client, err := setup.Setup()
	if err != nil {
		panic(err)
	}

	topicSubscribers := subscribers.RegisterSubscribers(client)

	for _, sub := range topicSubscribers {
		token := client.Subscribe(sub.GetTopic(), sub.GetQoS(), sub.Handler)
		token.Wait()
		fmt.Printf("Subscribed to %s \n", sub.GetTopic())
	}

	//token := client.Subscribe(topics.HomeLight, 1, nil)
	//token.Wait()

	fmt.Println("app is running")
	<-make(chan bool)
}

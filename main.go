package main

import (
	"fmt"
	"mqtt-client/setup"
	"time"
)

func main() {
	client, err := setup.Setup()
	if err != nil {
		panic(err)
	}

	topic := "home/light"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to %s \n", topic)

	num := 10
	for i := 0; i <= num; i++ {
		text := fmt.Sprintf("light %d", i)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	fmt.Println("app is running")
	<-make(chan bool)
}

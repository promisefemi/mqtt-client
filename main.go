package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("unable to load environment variables")
	}
	brokerIP := os.Getenv("MQTT_BROKER_IP")
	brokerPort := os.Getenv("MQTT_BROKER_PORT")
	clientID := os.Getenv("MQTT_CLIENT_ID")

	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%s", brokerIP, brokerPort))
	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(defaultPublishHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
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
func connectHandler(client mqtt.Client) {
	//fmt.Println(client)
	fmt.Println("Client is connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Fatalf("error: client is unable to connect - %s", err)
}

func defaultPublishHandler(client mqtt.Client, message mqtt.Message) {
	//fmt.Println(client)
	fmt.Printf("Message: %s received on topic: %s\n", message.Payload(), message.Topic())
}

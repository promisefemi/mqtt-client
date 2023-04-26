package subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mqtt-client/topics"
)

type HomeBedroomLightSubscriber struct{}

func NewHomeBedroomLightSubscriber() *HomeBedroomLightSubscriber {
	return &HomeBedroomLightSubscriber{}
}

func (h HomeBedroomLightSubscriber) GetTopic() string {
	return topics.HomeBedroomLight
}

func (h HomeBedroomLightSubscriber) GetQoS() byte {
	return 1
}

func (h HomeBedroomLightSubscriber) Handler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload(), message.Topic())
}

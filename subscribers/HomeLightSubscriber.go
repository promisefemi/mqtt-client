package subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"mqtt-client/topics"
)

type HomeLightSubscriber struct{}

func NewHomeLightSubscriber() *HomeLightSubscriber {
	return &HomeLightSubscriber{}
}

func (h HomeLightSubscriber) GetTopic() string {
	return topics.HomeLight
}

func (h HomeLightSubscriber) GetQoS() byte {
	return 1
}

func (h HomeLightSubscriber) Handler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload(), message.Topic())
}

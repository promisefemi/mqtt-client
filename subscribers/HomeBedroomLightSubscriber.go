package subscribers

import (
	"fmt"
	"github.com/eclipse/paho.golang/paho"
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

func (h HomeBedroomLightSubscriber) Handler(message *paho.Publish) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload, message.Topic)
}

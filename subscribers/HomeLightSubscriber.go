package subscribers

import (
	"fmt"
	"github.com/eclipse/paho.golang/paho"
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

func (h HomeLightSubscriber) Handler(message *paho.Publish) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload, message.Topic)
}

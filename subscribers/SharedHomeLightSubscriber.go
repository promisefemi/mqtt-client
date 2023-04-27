package subscribers

import (
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"mqtt-client/topics"
)

type SharedHomeLightSubscriber struct{}

func NewSharedHomeLightSubscriber() *SharedHomeLightSubscriber {
	return &SharedHomeLightSubscriber{}
}

func (h SharedHomeLightSubscriber) GetTopic() string {
	return fmt.Sprintf("$share/lightgroup/%s", topics.HomeLight)
}

func (h SharedHomeLightSubscriber) GetQoS() byte {
	return 1
}

func (h SharedHomeLightSubscriber) Handler(message *paho.Publish) {
	go func(message *paho.Publish) {
		fmt.Printf("Group Light: client 1 Message: %s -- topic: %s\n", message.Payload, message.Topic)
	}(message)
}

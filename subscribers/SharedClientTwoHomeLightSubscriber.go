package subscribers

import (
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"mqtt-client/topics"
)

type SharedClientTwoHomeLightSubscriber struct{}

func NewSharedClientTwoHomeLightSubscriber() *SharedClientTwoHomeLightSubscriber {
	return &SharedClientTwoHomeLightSubscriber{}
}

func (h SharedClientTwoHomeLightSubscriber) GetTopic() string {
	return fmt.Sprintf("$share/lightgroup/%s", topics.HomeLight)
}

func (h SharedClientTwoHomeLightSubscriber) GetQoS() byte {
	return 1
}

func (h SharedClientTwoHomeLightSubscriber) Handler(message *paho.Publish) {
	fmt.Printf("Group Light: client 2 Message: %s -- topic: %s\n", message.Payload, message.Topic)
}

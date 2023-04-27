package subscribers

import (
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"mqtt-client/topics"
)

type HomeThermostatSubscriber struct{}

func NewHomeThermostatSubscriber() *HomeThermostatSubscriber {
	return &HomeThermostatSubscriber{}
}

func (h HomeThermostatSubscriber) GetTopic() string {
	return topics.HomeThermostat
}

func (h HomeThermostatSubscriber) GetQoS() byte {
	return 1
}

func (h HomeThermostatSubscriber) Handler(message *paho.Publish) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload, message.Topic)
}

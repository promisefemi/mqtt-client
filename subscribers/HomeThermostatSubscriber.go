package subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

func (h HomeThermostatSubscriber) Handler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message: %s -- topic: %s\n", message.Payload(), message.Topic())
}

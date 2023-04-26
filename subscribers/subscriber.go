package subscribers

import mqtt "github.com/eclipse/paho.mqtt.golang"

type ISubscriber interface {
	GetTopic() string
	GetQoS() byte
	Handler(client mqtt.Client, message mqtt.Message)
}

type SystemSubscriber struct {
	subscribers []ISubscriber
}

func RegisterSubscribers(client mqtt.Client) []ISubscriber {
	allSubscribers := NewSystemSubscriber()

	allSubscribers.AddSubscriber(NewHomeLightSubscriber())
	allSubscribers.AddSubscriber(NewHomeThermostatSubscriber())
	allSubscribers.AddSubscriber(NewHomeBedroomLightSubscriber())

	return allSubscribers.ListSubscribers()
}

func NewSystemSubscriber() *SystemSubscriber {
	return &SystemSubscriber{
		subscribers: make([]ISubscriber, 0),
	}
}
func (s *SystemSubscriber) AddSubscriber(subscriber ISubscriber) {
	s.subscribers = append(s.subscribers, subscriber)
}

func (s *SystemSubscriber) ListSubscribers() []ISubscriber {
	return s.subscribers
}

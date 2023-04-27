package subscribers

import (
	"github.com/eclipse/paho.golang/paho"
)

type ISubscriber interface {
	GetTopic() string
	GetQoS() byte
	Handler(message *paho.Publish)
}

type SystemSubscriber struct {
	subscribers []ISubscriber
}

func RegisterSubscribers() []ISubscriber {
	allSubscribers := NewSystemSubscriber()
	allSubscribers.AddSubscriber(NewHomeLightSubscriber())
	allSubscribers.AddSubscriber(NewHomeThermostatSubscriber())
	allSubscribers.AddSubscriber(NewHomeBedroomLightSubscriber())
	allSubscribers.AddSubscriber(NewSharedClientTwoHomeLightSubscriber())
	allSubscribers.AddSubscriber(NewSharedHomeLightSubscriber())

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

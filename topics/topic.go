package topics

type Topic string

const (
	HomeLight      Topic = "home/light"
	HomeThermostat Topic = "home/thermostat"
)

func (t Topic) String() string {
	return t.String()
}

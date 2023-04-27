package setup

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

func Setup() (*paho.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("unable to load environment variables")
	}
	brokerIP := os.Getenv("MQTT_BROKER_IP")
	brokerPort := os.Getenv("MQTT_BROKER_PORT")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	//
	//options := mqtt.NewClientOptions()
	//options.AddBroker(fmt.Sprintf("tcp://%s:%s", brokerIP, brokerPort))
	//options.SetClientID(clientID)
	//options.SetDefaultPublishHandler(defaultPublishHandler)
	//options.OnConnect = connectHandler
	//options.OnConnectionLost = connectionLostHandler

	tcpServer := fmt.Sprintf("%s:%s", brokerIP, brokerPort)
	conn, err := net.Dial("tcp", tcpServer)
	if err != nil {
		return nil, err
	}
	options := paho.ClientConfig{
		Conn:   conn,
		Router: paho.NewStandardRouter(),
	}

	client := paho.NewClient(options)
	cp := &paho.Connect{
		KeepAlive:  30,
		ClientID:   clientID,
		CleanStart: true,
	}

	ca, err := client.Connect(context.Background(), cp)
	if err != nil {
		return nil, err
	}
	if ca.ReasonCode != 0 {
		return nil, fmt.Errorf("failed to connect to %s, Reason code: %d, Reason text: %s", tcpServer, ca.ReasonCode, ca.Properties.ReasonString)
	}
	fmt.Printf("Connected to %s\n", tcpServer)
	return client, nil
	//token := client.Connect()
	//
	//if token.Wait() && token.Error() != nil {
	//	return nil, token.Error()
	//}
	//return client, nil
}

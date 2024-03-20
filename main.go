package main

import (
	"fmt"
	"time"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type Data struct{
	Value string `json:"value"`
}

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	protocol := os.Getenv("PROTOCOL")
	url := os.Getenv("URL")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	topic := os.Getenv("TOPIC")
	
	opts := MQTT.NewClientOptions().AddBroker(protocol + "://" + url)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := MQTT.NewClient(opts)

	token := client.Connect()
	if token.Wait() && token.Error() != nil{
		panic(token.Error())
	}
	defer client.Disconnect(250)

	for{
		data := Data{Value: "Ini adalah data"}
		sendToken := client.Publish(topic, 0, false, data.Value)
		sendToken.Wait()

		fmt.Println("Data berhasil dikirim")
		time.Sleep(10 * time.Second)
	}
}

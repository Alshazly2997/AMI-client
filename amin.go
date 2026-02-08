package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/heltonmarx/goami/ami"
)

var (
	//You have to set your AMI username, secret, and IP address here
	username = flag.String("username", "username", "AMI username")
	secret   = flag.String("secret", "secret", "AMI secret")
	host     = flag.String("host", "172.0.0.1:5038", "AMI host address")
)

func main() {
	flag.Parse()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	socket, err := ami.NewSocket(*host)
	if err != nil {
		log.Fatalf("socket error: %v\n", err)
	}
	if _, err := ami.Connect(socket); err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	//Login
	uuid := fmt.Sprintf("%d", time.Now().UnixNano())

	if _, err := ami.Login(socket, *username, *secret, "On", uuid); err != nil {
		log.Fatalf("login error: %v\n", err)
	}
	fmt.Printf("login ok!\n")

	//Listen for events
	fmt.Printf("Listening for events... (Press Ctrl+C to stop)\n")

	//open a text file to register events
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file open error: %v\n", err)
	}
	defer file.Close()

	for {

		event, err := socket.Recv()
		if err != nil {
			log.Printf("recv error: %v\n", err)
			continue
		}

		fmt.Printf("Received event: %s\n", event)
		_, err = file.WriteString(event + "\n")
		if err != nil {
			fmt.Printf("Failed to write to log file: %v\n", err)
		}
	}
}

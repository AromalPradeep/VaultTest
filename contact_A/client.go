// socket-client project main.go
package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func client() {

	// establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}

	// send some data
	_, err = connection.Write([]byte("Hello Server! Greetings from contact_A."))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	println("Message Sent!")

	// reading response
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Recieved Response from Contact B: ", string(buffer[:mLen]))
	defer connection.Close()
}

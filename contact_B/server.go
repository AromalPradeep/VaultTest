// socket-server project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVERA_HOST = "localhost"
	SERVERA_PORT = "9988"
	SERVERA_TYPE = "tcp"
)

func server() {
	fmt.Println("Server B Running...")
	server, err := net.Listen(SERVERA_TYPE, SERVERA_HOST+":"+SERVERA_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected.")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	connection.Close()
}

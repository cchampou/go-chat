package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Conn net.Conn
	Name string
}

func main() {
	server := createServer()
	pipe, errListen := net.Listen(server.Network, server.Port)
	handleError(errListen, true)
	fmt.Println("server started")

	var clients []Client

	for {
		conn, errAccept := pipe.Accept()
		handleError(errAccept, true)

		fmt.Println("New client from", conn.RemoteAddr())

		go func() {
			buf := bufio.NewReader(conn)
			name, nameErr := buf.ReadString('\n')
			handleError(nameErr, false)
			currentClient := Client{Name: strings.TrimSuffix(name, "\n"), Conn: conn}
			clients = append(clients, currentClient)
			println("Client picked username " + currentClient.Name)
			for {
				msg, msgErr := buf.ReadString('\n')
				if msgErr != nil {
					println(currentClient.Name + " disconnected")
					break
				}
				print("Message from " + currentClient.Name + ": " + msg)

				for _, c := range clients {
					if c.Name != currentClient.Name {
						_, _ = c.Conn.Write([]byte(currentClient.Name + ": " + msg))
					}
				}
			}
		}()
	}
}

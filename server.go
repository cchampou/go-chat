package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	Conn net.Conn
	Name string
}

func log(msg string) {
	start := time.Now()
	fmt.Println(start.String() + " " + msg)
}

func main() {
	server := createServer()
	pipe, errListen := net.Listen(server.Network, server.Port)
	handleError(errListen, true)
	log("Server started")

	var clients []Client

	for {
		conn, errAccept := pipe.Accept()
		handleError(errAccept, true)

		log("New client from " + conn.RemoteAddr().String())

		go func() {
			buf := bufio.NewReader(conn)
			name, nameErr := buf.ReadString('\n')
			handleError(nameErr, false)
			currentClient := Client{Name: strings.TrimSuffix(name, "\n"), Conn: conn}
			clients = append(clients, currentClient)
			log("Client picked username " + currentClient.Name)
			for {
				msg, msgErr := buf.ReadString('\n')
				if msgErr != nil {
					log(currentClient.Name + " disconnected")
					break
				}
				log("Message from " + currentClient.Name + ": " + strings.TrimSuffix(msg, "\n"))

				for _, c := range clients {
					if c.Name != currentClient.Name {
						_, _ = c.Conn.Write([]byte(currentClient.Name + ": " + msg))
					}
				}
			}
		}()
	}
}

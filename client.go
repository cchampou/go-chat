package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	server := createServer()

	conn, errDial := net.Dial(server.Network, server.Port)
	handleError(errDial, true)
	fmt.Println("Client connected")

	wg.Add(2)

	localBuf := bufio.NewReader(os.Stdin)
	remoteBuf := bufio.NewReader(conn)

	// Ask for username
	print("Pick a username: ")
	name, errInput := localBuf.ReadString('\n')
	handleError(errInput, true)
	_, _ = conn.Write([]byte(name))

	// Goroutine for user input
	go func() {
		defer wg.Done()
		for {
			msg, errMsg := localBuf.ReadString('\n')
			handleError(errMsg, true)
			_, _ = conn.Write([]byte(msg))
		}
	}()

	// Goroutine to print server messages
	go func() {
		defer wg.Done()
		for {
			serverMsg, errServerMsg := remoteBuf.ReadString('\n')
			handleError(errServerMsg, true)
			print(serverMsg)
		}
	}()

	wg.Wait()
}

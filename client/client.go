package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"cchampou.me/utils"
)

func main() {

	var wg sync.WaitGroup

	server := utils.CreateServer()

	conn, errDial := net.Dial(server.Network, server.Port)
	utils.HandleError(errDial, true)
	fmt.Println("Client connected")

	wg.Add(2)

	localBuf := bufio.NewReader(os.Stdin)
	remoteBuf := bufio.NewReader(conn)

	// Ask for username
	print("Pick a username: ")
	name, errInput := localBuf.ReadString('\n')
	utils.HandleError(errInput, true)
	_, _ = conn.Write([]byte(name))

	// Goroutine for user input
	go func() {
		defer wg.Done()
		for {
			msg, errMsg := localBuf.ReadString('\n')
			utils.HandleError(errMsg, true)
			_, _ = conn.Write([]byte(msg))
		}
	}()

	// Goroutine to print server messages
	go func() {
		defer wg.Done()
		for {
			serverMsg, errServerMsg := remoteBuf.ReadString('\n')
			utils.HandleError(errServerMsg, true)
			print(serverMsg)
		}
	}()

	wg.Wait()
}

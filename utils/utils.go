package utils

import (
	"fmt"
	"os"
)

type Server struct {
	Port    string
	Network string
}

func CreateServer() Server {
	server := Server{Port: ":5000", Network: "tcp"}
	return server
}

func HandleError(err error, fatal bool) {
	if err != nil {
		fmt.Println(err)
		if fatal {
			os.Exit(1)
		}
	}
}

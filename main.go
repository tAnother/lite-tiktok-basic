package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/config"
)

func main() {
	var port string
	var serverIP string

	// Get server IP from user input，default 127.0.0.1
	for {
		fmt.Print("Enter server ip_address (press Enter to use default 127.0.0.1): ")
		fmt.Scanln(&serverIP)

		if serverIP == "" {
			serverIP = "127.0.0.1"
			break
		}
		if net.ParseIP(serverIP) != nil {
			break
		}

		fmt.Println("invalid ip address, please retry.")
	}

	// Get server port from user input, default 8080
	for {
		fmt.Print("Enter port (press Enter to use default 8080): ")
		fmt.Scanln(&port)

		if port == "" {
			port = "8080"
			break
		}
		if p, err := strconv.Atoi(port); err == nil && p > 0 && p < 65536 {
			break
		}

		fmt.Println("invalid port number, please retry.")
	}

	config.SetIPAndPort(serverIP, port)
	// go service.RunMessageServer()

	r := gin.Default()

	config.SqlInit()
	config.RedisInit()
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

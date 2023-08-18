package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/tools"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
)

func main() {
	var port string
	var serverIP string

	// 循环直到获取一个有效的IP
	for {
		fmt.Print("please port the server ip_address: ")
		fmt.Scanln(&serverIP)

		if net.ParseIP(serverIP) != nil {
			break
		}

		fmt.Println("invalid ip address, please re-port.")
	}

	// 循环直到获取一个有效的端口
	for {
		fmt.Print("please port the server port: ")
		fmt.Scanln(&port)

		p, err := strconv.Atoi(port)
		if err == nil && p > 0 && p < 65536 {

			break
		}

		fmt.Println("invalid port number, please re-port.")
	}

	tools.ServerSetting(serverIP, port)
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)
	tools.SqlInit()
	tools.RedisInit()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

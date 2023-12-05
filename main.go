package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("check", func(c *gin.Context) {
		ipAddress := c.ClientIP()

		fmt.Println(ipAddress)

		c.JSON(200, ipAddress)
	})

	server.Run(":7000")
}
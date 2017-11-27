package main

import (
	"github.com/gin-gonic/gin"
	fcm "github.com/mochadwi/push-notif/controller"
)

func main() {
	router := gin.Default()
	var baseURL = "/api/v1"

	v1 := router.Group(baseURL)
	{
		v1.POST("/push-notif", fcm.SendGMToClient)
	}

	router.Run(":8080")
}

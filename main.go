package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	fcm "gitlab.com/nobackend-repo/push-notif-service/controllers"
	Index "gitlab.com/nobackend-repo/push-notif-service/views"
)

func main() {
	var baseURL = "/api/v1"
	engine := gin.Default()
	engine.RedirectTrailingSlash = false
	// engine.Use(utils.Middleware)

	v1 := engine.Group(baseURL)
	{
		// single-device push notif
		v1.POST("/push-notif", fcm.SendGMToClient)

		// multiple devices push notif
		v1.POST("/device/:any/notifications", fcm.SendGMToClients)
	}

	engine.NoRoute(func(c *gin.Context) {
		var response = &Index.DefaultResponseFormat{
			RequestID: uuid.NewV4().String(),
			Now:       time.Now().Unix(),
		}

		response.Code = strconv.Itoa(http.StatusNotFound) + "00"
		response.Message = "Service resource not found"

		c.JSON(http.StatusNotFound, response)
	})

	engine.Run(":8080")
}

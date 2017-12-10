package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/nobackend-repo/push-notif-service/models"
	db "gitlab.com/nobackend-repo/push-notif-service/utils"
)

// CreateNotifier register a notifier
func CreateNotifier(c *gin.Context) {
	var notifier models.NotifierItem
	c.BindJSON(&notifier)
	db.Mgr.AddNotifier(&notifier)
	c.JSON(200, notifier)
}

// GetNotifiers all available list
func GetAllNotifier(c *gin.Context) {
	notifier := []models.NotifierItem{}

	if err := db.Mgr.ShowAllNotifier(&notifier); err != nil {
		c.AbortWithStatus(404)
	} else {
		fmt.Print("[notifier] results: ")
		fmt.Print(notifier)
		c.JSON(200, notifier)
	}
}

// GetNotifier list notifier based on name
func GetNotifier(c *gin.Context) {
	name := c.Params.ByName("name")
	var notifier models.NotifierItem
	if err := db.Mgr.ShowNotifier(name, &notifier); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(404, c.BindJSON(&err))
	} else {
		c.JSON(200, notifier)
	}
}

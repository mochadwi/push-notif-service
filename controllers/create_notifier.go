package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	models "gitlab.com/nobackend-repo/push-notif-service/models"
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
func GetNotifiers(c *gin.Context) {
	var notifier []models.NotifierItem
	if err := db.Mgr.ShowNotifier(notifier); err != nil {
		c.AbortWithStatus(404)
		fmt.Print("GetNotifiers: ")
		fmt.Print(err)
	} else {
		c.JSON(200, notifier)
	}
}

// GetNotifier list notifier based on name
// func GetNotifier(c *gin.Context) {
// 	name := c.Params.ByName("name")
// 	var notifier models.NotifierItem
// 	if err := db.Where("name = ?", name).First(&notifier).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 		c.JSON(404, c.BindJSON(&err))
// 	} else {
// 		c.JSON(200, notifier)
// 	}
// }

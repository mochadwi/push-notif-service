package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/nobackend-repo/push-notif-service/models"
	db "gitlab.com/nobackend-repo/push-notif-service/utils"
	Index "gitlab.com/nobackend-repo/push-notif-service/views"
)

// CreateNotifier register a notifier
func CreateNotifier(c *gin.Context) {
	var notifier models.NotifierItem
	c.BindJSON(&notifier)

	err := db.Mgr.AddNotifier(&notifier)

	var response = &Index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = strconv.Itoa(http.StatusBadRequest) + "01"
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = strconv.Itoa(http.StatusOK) + "01"
		response.Message = "OK"
		response.Data = notifier

		c.JSON(http.StatusCreated, response)
	}
}

// GetAllNotifier all available list
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

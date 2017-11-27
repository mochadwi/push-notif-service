package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gcm "github.com/google/go-gcm"
	uuid "github.com/mochadwi/push-notif/utils"
)

// SendGMToClient is a function that will push a message to client
func SendGMToClient(c *gin.Context) {
	serverKey := "AAAAW65R3Ag:APA91bEkID7zMUqEWVEh-spXN1NkD6jCPE4mxnK23ocKrCJWogV5RvIuPOO98JhwjbwqA4C5zXEkpSmkzH-rQJmQLKxtsgy9LzWp8LtYoln7OBEbEFVPQgBmVYeE1_qLCO1zZ7No5HeG"
	var msg gcm.HttpMessage
	data := map[string]interface{}{"message": c.Query("message")}
	regIDs := []string{c.Query("client_token")}
	msg.RegistrationIds = regIDs
	msg.Data = data
	response, err := gcm.SendHttp(serverKey, msg)

	t := time.Now()
	uuid, errUUID := uuid.NewUUID()
	if errUUID != nil {
		fmt.Printf("error: %v\n", err)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": uuid,
			"now":       t.Format("2006/01/02 15:04:05"),
			"code":      strconv.Itoa(http.StatusBadRequest) + "02",
			"message":   err.Error(),
			"data":      "[]",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"requestID": uuid,
			"now":       t.Format("2006/01/02 15:04:05"),
			"code":      strconv.Itoa(http.StatusOK) + "01",
			"message":   response.Error,
			"data":      response.Results,
		})
	}
}

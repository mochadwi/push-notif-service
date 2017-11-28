package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gcm "github.com/google/go-gcm"
	uuid "github.com/satori/go.uuid"
	Index "gitlab.com/nobackend-repo/push-notif-service/views"
)

// SendGMToClients will push a message to multi clients
func SendGMToClients(c *gin.Context) {
	var msg gcm.HttpMessage
	serverKey := c.Query("server_key")
	data := map[string]interface{}{"message": c.Query("message")}
	regIDs := []string{c.Query("client_token")}
	msg.RegistrationIds = regIDs
	msg.Data = data
	gcmResp, err := gcm.SendHttp(serverKey, msg)

	var response = &Index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = strconv.Itoa(http.StatusBadRequest) + "02"
		response.Message = err.Error()
		response.Data = "[]"

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = strconv.Itoa(http.StatusOK) + "01"
		response.Message = gcmResp.Error
		response.Data = gcmResp.Results

		c.JSON(http.StatusOK, response)
	}
}

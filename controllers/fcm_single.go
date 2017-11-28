package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gcm "github.com/google/go-gcm"
	uuid "github.com/satori/go.uuid"
	Index "gitlab.com/nobackend-repo/push-notif-service/views"
)

// SendGMToClient will push a message to single client
func SendGMToClient(c *gin.Context) {
	var msg gcm.HttpMessage
	serverKey := c.Query("server_key")
	data := map[string]interface{}{
		"title": c.Query("title"),
		"body":  c.Query("body")}

	regIDs := []string{c.Query("client_token")}
	msg.RegistrationIds = regIDs
	msg.Data = data
	gcmResp, err := gcm.SendHttp(serverKey, msg)
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)

	var response = &Index.DefaultResponseFormat{
		RequestID: uuid.NewV4().String(),
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = strconv.Itoa(http.StatusBadRequest) + "02"
		response.Message = err.Error()
		response.Data = "[]"

		c.JSON(http.StatusBadRequest, response)
		fmt.Printf("Body is : %s \n", string(x))
	} else {
		response.Code = strconv.Itoa(http.StatusOK) + "01"
		response.Message = gcmResp.Error
		response.Data = gcmResp.Results

		c.JSON(http.StatusOK, response)
		fmt.Printf("Body is : %s \n", string(x))
	}
}

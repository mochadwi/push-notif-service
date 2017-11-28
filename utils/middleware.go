package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Middleware function will validate json scheme
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Will be doing json schema validation here

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		fmt.Printf("%s \n", string(x))

		fmt.Println("I am a middleware for json schema validation")

		c.Next()
		return
	}
}

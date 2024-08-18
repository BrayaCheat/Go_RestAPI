package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Get the start time
		startTime := time.Now()

		//Process time
		c.Next()

		//Get the end time
		endTime := time.Now()

		//Log the request method, path, and duration
		latency := endTime.Sub(startTime)

		fmt.Printf("Logger => Method: %s | Endpoint: %s | Duration: %v\n", c.Request.Method, c.Request.URL.Path, latency)
	}
}

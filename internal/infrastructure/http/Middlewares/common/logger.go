package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Stop timer
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// Get status
		statusCode := c.Writer.Status()

		// Get client IP
		clientIP := c.ClientIP()

		// Get request method
		reqMethod := c.Request.Method

		// Get request URI
		reqURI := c.Request.RequestURI

		// Get user agent
		userAgent := c.Request.UserAgent()

		// Log only when status code is greater than or equal to 400
		if statusCode >= 400 {
			logrus.WithFields(logrus.Fields{
				"status_code":  statusCode,
				"latency_time": latencyTime,
				"client_ip":    clientIP,
				"req_method":   reqMethod,
				"req_uri":      reqURI,
				"user_agent":   userAgent,
			}).Error("HTTP request error")
		} else {
			logrus.WithFields(logrus.Fields{
				"status_code":  statusCode,
				"latency_time": latencyTime,
				"client_ip":    clientIP,
				"req_method":   reqMethod,
				"req_uri":      reqURI,
				"user_agent":   userAgent,
			}).Info("HTTP request")
		}
	}
}

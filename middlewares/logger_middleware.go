package middlewares

import (
	"bytes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ResponseWriter is a custom response writer to capture the response body
type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// LoggerMiddleware logs the details of each request and response
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New()
		startTime := time.Now()
		// Create a custom response writer
		rw := &ResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rw

		// Process the request
		c.Next()

		// Calculate the time taken to process the request
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Log the request and response details
		log.Printf("| ID: %s | Request: %s %s | Status: %d | Latency: %s | Response: %s",
			id,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			rw.body.String(),
		)
	}
}

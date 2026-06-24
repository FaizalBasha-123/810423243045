package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		duration := time.Since(start)
		log.Printf("[%s] %s %s took %v",
			context.Request.Method,
			context.Request.URL.Path,
			context.ClientIP(),
			duration,
		)
	}
}

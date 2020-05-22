package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger(log logrus.FieldLogger) gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		stop := time.Since(start)

		hostName, hostError := os.Hostname()

		if hostError != nil {
			hostName = "Unknown host"
		}

		entry := log.WithFields(logrus.Fields{
			"hostname":        hostName,
			"statusCode":      context.Writer.Status(),
			"latency":         int(math.Ceil(float64(stop.Nanoseconds())) / 1000000),
			"clientIP":        context.ClientIP(),
			"method":          context.Request.Method,
			"path":            context.Request.URL.Path,
			"referer":         context.Request.Referer(),
			"dataLength":      math.Max(float64(context.Writer.Size()), 0),
			"clientUserAgent": context.Request.UserAgent(),
		})

		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			message := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s %s\" (%d ms)",
				context.ClientIP(),
				hostName,
				time.Now().Format("02 Jan 2006 15:04:05 -0700"),
				context.Request.Method,
				context.Request.URL.Path,
				context.Writer.Status(),
				context.Writer.Size(),
				context.Request.Referer(),
				context.Request.UserAgent(),
				int(math.Ceil(float64(stop.Nanoseconds()))/1000000),
			)

			if context.Writer.Status() > 499 {
				entry.Error(message)
			} else if context.Writer.Status() > 399 {
				entry.Warn(message)
			} else {
				entry.Info(message)
			}
		}
	}
}

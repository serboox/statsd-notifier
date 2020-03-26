package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
)

// Logger sends request data to the output stream
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidType, err := uuid.NewV4()
		if err != nil {
			log.Error(err)
		}

		requestID := uuidType.String()

		c.Set("request_uuid", requestID)
		c.Set("logger", log.WithField("request_uuid", requestID))

		start := time.Now().UTC()

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		var fields log.Fields = make(map[string]interface{})
		fields["date"] = start.Format("2006-01-02")
		period := fmt.Sprintf(
			"(%s)->(%s)",
			start.Format("15:04:05.999"),
			end.Format("15:04:05.999"),
		)
		fields["time"] = period
		fields["uuid"] = requestID

		e, exists := c.Get("api_error")
		if exists {
			fields["error"] = e.(error).Error()
		}

		ginErr := c.Errors.String()
		if ginErr != "" {
			log.Error(ginErr)
		}

		log.WithFields(fields).Infof(
			"GIN: %s %s %s code=%d ip=%s",
			c.Request.Method,
			c.Request.URL.Path,
			fmt.Sprint(latency),
			c.Writer.Status(),
			c.ClientIP(),
		)
	}
}

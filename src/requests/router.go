package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/utils"
)

// SetupRouter Implements mapping of the path to the function
func SetupRouter() (r *gin.Engine) {
	r = gin.New()
	r.POST("/", rootPost)
	r.NoRoute(noRoute)

	r.Use(utils.Logger())

	return
}

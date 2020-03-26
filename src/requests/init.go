package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/configs"
	"github.com/serboox/statsd-notifier/src/consts"
	"github.com/serboox/statsd-notifier/src/utils"
)

// SetupRouter Implements mapping of the path to the function
func SetupRouter(ctx *configs.Context) (r *gin.Engine) {
	r = gin.New()

	// Detailed requests logging
	r.Use(utils.Logger())
	// Recovers if panic
	r.Use(gin.Recovery())
	r.Use(addContext(ctx))

	r.POST("/", rootPost)
	r.NoRoute(noRoute)

	return
}

func addContext(ctx *configs.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(consts.FieldContext, ctx)
	}
}

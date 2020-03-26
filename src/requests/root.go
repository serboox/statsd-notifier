package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/configs"
	"github.com/serboox/statsd-notifier/src/consts"
)

func rootPost(c *gin.Context) {
	ctx := c.MustGet(consts.FieldContext).(*configs.Context)
	//b.reqUUID = b.ginCtx.MustGet(commConsts.FieldRequestUUID).(string)

	// Create a new StatsD connection
	// host := "localhost"
	// port := 8125

	// client := statsd.New(host, port)
	// client.Increment("production.fqdn.statsd.post.request.counter")
	ctx.StatsD.Increment("production.fqdn.statsd.post.request.counter")

	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

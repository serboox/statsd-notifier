package requests

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/configs"
	"github.com/serboox/statsd-notifier/src/consts"
	"github.com/serboox/statsd-notifier/src/exceptions"
)

// GeoJSON contains information on geo location
// e.g. {"geo":{"CityName":"New York City","ContinentCode":"NA","CountryIsoCode":"US"}}
type GeoJSON struct {
	Geo struct {
		CityName       string `json:"CityName" binding:"required"`
		ContinentCode  string `json:"ContinentCode" binding:"required"`
		CountryIsoCode string `json:"CountryIsoCode" binding:"required"`
	} `json:"geo"`
}

func rootPost(c *gin.Context) {
	ctx := c.MustGet(consts.FieldContext).(*configs.Context)

	if count, ok := c.GetQuery(consts.FieldCount); ok {
		log.Debugf("Count number: %s\n", count)
	}

	geoJSON := new(GeoJSON)
	if err := c.BindJSON(geoJSON); err != nil {
		log.Debugln("Json unmarshal error: ", err)
		exceptions.BadRequest(err.Error()).Response(c)

		return
	}

	log.Debugf("Request JSON: %+v", geoJSON)

	go ctx.StatsD.Increment("production.fqdn.statsd.post.request.counter")

	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

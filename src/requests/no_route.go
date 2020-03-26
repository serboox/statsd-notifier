package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func noRoute(c *gin.Context) {
	c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

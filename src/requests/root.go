package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootPost(c *gin.Context) {
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

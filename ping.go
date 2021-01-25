package ginping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping just replies "OK"
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

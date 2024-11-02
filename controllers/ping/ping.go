package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Tags ping
// @Description ping to service
// @ID ping
// @Produce  text/plain; charset=utf-8
// @Success 200 {string} string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

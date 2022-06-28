package bluecore

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldHandler(c *gin.Context) {
	world := c.Param("world")

	if world == "" {
		c.String(http.StatusBadRequest, "/hello/{world} contained no world data")
	} else if len(world) > 10 {
		c.String(http.StatusBadRequest, "/hello/{world} contained too long of data: %s", world)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Hello": world,
		})
	}
}

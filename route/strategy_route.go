package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/littlecheny/chenzhangtest/services"
)

func newStrategyRoute(publicRoute *gin.RouterGroup, taskManager *services.TaskManager) {
	publicRoute.POST("/strategy", func(c *gin.Context) {
		var req struct {
			Algo string `json:"algo"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		taskManager.ChangeStrategy(req.Algo)
		fmt.Printf("strategy changed to %s\n", req.Algo)
		c.JSON(http.StatusOK, gin.H{"message": "strategy changed"})
	})
}

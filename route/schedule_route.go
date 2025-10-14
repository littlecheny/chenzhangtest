package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/littlecheny/chenzhangtest/domain"
	"github.com/littlecheny/chenzhangtest/services"
)

func newScheduleRoute(r *gin.RouterGroup, taskManager *services.TaskManager) {

	// 解析请求，创建tasks[]
	r.POST("/submitasks", func(c *gin.Context) {
		var addTasks []domain.Task
		if err := c.ShouldBindJSON(&addTasks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			if err := taskManager.AddTasks(c.GetString("userID"), addTasks); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Tasks added successfully"})
		}
	})
}

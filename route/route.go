package route

import (
	"github.com/gin-gonic/gin"
	"github.com/littlecheny/chenzhangtest/services"
)

func Setup(r *gin.Engine, taskManager *services.TaskManager) {
	publicRoute := r.Group("")
	newScheduleRoute(publicRoute, taskManager)
	newStrategyRoute(publicRoute, taskManager)
}

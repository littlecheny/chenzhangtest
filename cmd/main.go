package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/littlecheny/chenzhangtest/route"
	"github.com/littlecheny/chenzhangtest/services"
)

func main() {
	r := gin.Default()
	taskManager := services.NewTaskManager()

	route.Setup(r, taskManager)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go taskManager.ScheduleLoop(ctx, time.Second, "FIFO")

	r.Run(":8080")
}

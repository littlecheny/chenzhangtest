package domain

// 任务
type Task struct {
	UserID      string // 用户ID
	Index       int    // 任务索引
	ArrivalTime int    // 到达时间
	BurstTime   int    // 执行时间片
}

type TaskManager interface {
	AddTasks(tasks []Task) error
	Snapshot() ([]Task, error)
	ScheduleNow(algo string) (SchedulerState, error)
	Clear(userID string) error
}

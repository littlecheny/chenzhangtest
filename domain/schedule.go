package domain

import (
	"time"
)

type Schedule interface {
	FIFOSchedule(t time.Time, tasks []Task) SchedulerState
	SRFTSchedule(t time.Time, tasks []Task) SchedulerState
	MergeTask(tasks []Task, stats SchedulerState) []Task
}

// 调度输出状态
type SchedulerState struct {
	Time time.Time
	// 执行任务索引 (index)
	ScheduledIndexes []string
	//使用一个辅助结构体来明确任务和剩余时间的关系
	RemainingTimes []TaskRemainingTime // 更清晰
}

// 记录被调度任务的剩余时间
type TaskRemainingTime struct {
	Index         string // 任务索引
	RemainingTime int    // 剩余执行时间片
	// 还可以加入一个字段来存储计算过程的描述，例如 "任务 1:4-3=1" [cite: 15]
	Description string
}

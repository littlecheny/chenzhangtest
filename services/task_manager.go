package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/littlecheny/chenzhangtest/domain"
)

type TaskManager struct {
	mu              sync.RWMutex
	tasks           []domain.Task
	scheduleService domain.Schedule
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:           make([]domain.Task, 0),
		scheduleService: NewFIFOScheduleService(),
	}
}

// AddTasks 将任务追加到用户维度的任务列表中，设置 UserID，并做切片拷贝以避免外部修改内部状态
func (m *TaskManager) AddTasks(userID string, tasks []domain.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range tasks {
		// 绑定用户ID
		m.tasks = append(m.tasks, t)
	}
	return nil
}

// Snapshot 返回用户当前任务的深拷贝快照，避免并发读写冲突
func (m *TaskManager) Snapshot() ([]domain.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	src := m.tasks
	snapshot := make([]domain.Task, len(src))
	copy(snapshot, src)
	return snapshot, nil
}

// ScheduleNow 基于当前快照执行调度，并将结果合并回用户任务列表（加写锁）
func (m *TaskManager) ScheduleNow(t time.Time, algo string) (domain.SchedulerState, error) {
	// TODO: 根据 algo 选择具体算法；当前使用默认 scheduleService。
	_ = algo

	// 读取快照不阻塞写操作
	tasks, _ := m.Snapshot()

	if len(tasks) == 0 {
		return domain.SchedulerState{Time: t}, nil
	}

	stats := m.scheduleService.Schedule(t, tasks)

	// 合并结果需要写锁，避免与其他写操作竞争
	m.mu.Lock()
	updated := m.scheduleService.MergeTask(m.tasks, stats)
	m.tasks = updated
	m.mu.Unlock()
	return stats, nil
}

// Clear 清空某个用户的任务
func (m *TaskManager) Clear(userID string) error {
	return nil
}

// ScheduleLoop 后台持续执行调度，以固定间隔遍历所有用户并触发一次调度
func (m *TaskManager) ScheduleLoop(ctx context.Context, interval time.Duration, algo string) {
	if interval <= 0 {
		interval = time.Second
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ticker.C:
			stats, _ := m.ScheduleNow(t, algo)
			// 格式化输出到秒，去除毫秒
			fmt.Printf("Time: %s, Scheduled: %v, Remaining: %v\n",
				stats.Time.Format("15:04:05"),
				stats.ScheduledIndexes, stats.RemainingTimes)

		}
	}
}

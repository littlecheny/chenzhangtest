package services

import (
	"slices"
	"sort"
	"strconv"
	"time"

	"github.com/littlecheny/chenzhangtest/domain"
)

type SRFTScheduleService struct {
}

func NewSRFTScheduleService() domain.Schedule {
	return &SRFTScheduleService{}
}

func (s *SRFTScheduleService) Schedule(t time.Time, tasks []domain.Task) domain.SchedulerState {
	resource := 5
	stats := domain.SchedulerState{
		Time: t,
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].BurstTime < tasks[j].BurstTime
	})
	for _, task := range tasks {
		if 0 < resource {
			description1 := strconv.Itoa(task.BurstTime) + "-" + strconv.Itoa(min(resource, task.BurstTime)) + "=" + strconv.Itoa(task.BurstTime-min(resource, task.BurstTime))
			resource -= min(resource, task.BurstTime)
			task.BurstTime -= min(resource, task.BurstTime)
			stats.ScheduledIndexes = append(stats.ScheduledIndexes, task.UserID+strconv.Itoa(task.Index))
			stats.RemainingTimes = append(stats.RemainingTimes, domain.TaskRemainingTime{
				Index:         task.UserID + strconv.Itoa(task.Index),
				RemainingTime: task.BurstTime,
				Description:   "任务 " + task.UserID + strconv.Itoa(task.Index) + ":" + description1,
			})
			if resource <= 0 {
				break
			}
		}
	}
	return stats
}

func (s *SRFTScheduleService) MergeTask(tasks []domain.Task, stats domain.SchedulerState) []domain.Task {
	for _, rt := range stats.RemainingTimes {
		if rt.RemainingTime == 0 {
			for i, task := range tasks {
				if task.UserID+strconv.Itoa(task.Index) == rt.Index {
					tasks = slices.Delete(tasks, i, i+1)
					break
				}
			}
		} else {
			for i, task := range tasks {
				if task.UserID+strconv.Itoa(task.Index) == rt.Index {
					tasks[i].BurstTime = rt.RemainingTime
					break
				}
			}
		}
	}
	return tasks
}

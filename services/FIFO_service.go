package services

import (
	"slices"
	"strconv"
	"time"

	"github.com/littlecheny/chenzhangtest/domain"
)

type ScheduleService struct {
}

func NewFIFOScheduleService() domain.Schedule {
	return &ScheduleService{}
}

func (s *ScheduleService) FIFOSchedule(t time.Time, tasks []domain.Task) domain.SchedulerState {
	resource := 5
	stats := domain.SchedulerState{
		Time: t,
	}
	for _, task := range tasks {
		if 0 < resource {
			consume := min(resource, task.BurstTime)
			description1 := strconv.Itoa(task.BurstTime) + "-" + strconv.Itoa(consume) + "=" + strconv.Itoa(task.BurstTime-consume)
			resource -= consume
			task.BurstTime -= consume
			stats.ScheduledIndexes = append(stats.ScheduledIndexes, task.UserID+"-"+strconv.Itoa(task.Index))
			stats.RemainingTimes = append(stats.RemainingTimes, domain.TaskRemainingTime{
				Index:         task.UserID + "-" + strconv.Itoa(task.Index),
				RemainingTime: task.BurstTime,
				Description:   "任务 " + task.UserID + "-" + strconv.Itoa(task.Index) + ":" + description1,
			})
			if resource <= 0 {
				break
			}
		}
	}
	return stats
}

func (s *ScheduleService) SRFTSchedule(t time.Time, tasks []domain.Task) domain.SchedulerState {
	resource := 5
	stats := domain.SchedulerState{
		Time: t,
	}
	slices.SortFunc(tasks, func(a, b domain.Task) int {
		return a.BurstTime - b.BurstTime
	})
	for _, task := range tasks {
		if 0 < resource {
			consume := min(resource, task.BurstTime)
			description1 := strconv.Itoa(task.BurstTime) + "-" + strconv.Itoa(consume) + "=" + strconv.Itoa(task.BurstTime-consume)
			resource -= consume
			task.BurstTime -= consume
			stats.ScheduledIndexes = append(stats.ScheduledIndexes, task.UserID+"-"+strconv.Itoa(task.Index))
			stats.RemainingTimes = append(stats.RemainingTimes, domain.TaskRemainingTime{
				Index:         task.UserID + "-" + strconv.Itoa(task.Index),
				RemainingTime: task.BurstTime,
				Description:   "任务 " + task.UserID + "-" + strconv.Itoa(task.Index) + ":" + description1,
			})
			if resource <= 0 {
				break
			}
		}
	}
	return stats
}

func (s *ScheduleService) MergeTask(tasks []domain.Task, stats domain.SchedulerState) []domain.Task {
	for _, rt := range stats.RemainingTimes {
		if rt.RemainingTime == 0 {
			for i, task := range tasks {
				if task.UserID+"-"+strconv.Itoa(task.Index) == rt.Index {
					tasks = slices.Delete(tasks, i, i+1)
					break
				}
			}
		} else {
			for i, task := range tasks {
				if task.UserID+"-"+strconv.Itoa(task.Index) == rt.Index {
					tasks[i].BurstTime = rt.RemainingTime
					break
				}
			}
		}
	}
	return tasks
}

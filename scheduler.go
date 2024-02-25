package main

import (
	"fmt"
	"time"
)

type Task interface {
    Execute()
}

type PrintTask struct {
    Message string
}

func (p PrintTask) Execute() {
    fmt.Println(p.Message)
}

type TaskScheduler struct {
    TaskQueue chan Task
    Interval time.Duration
}

func NewTaskScheduler(interval time.Duration) *TaskScheduler {
    return &TaskScheduler {
        TaskQueue: make(chan Task),
        Interval: interval,
    }
}

func (s *TaskScheduler) Start() {
    go func() {
        ticker := time.NewTicker(s.Interval)

        for {
            select {
            case task := <-s.TaskQueue:
                task.Execute()
            case <-ticker.C:
                for task := range s.TaskQueue {
                    task.Execute()
                }
            }
        }
    }()
}

func (s *TaskScheduler) ScheduleOnce(dur time.Duration, task Task) {
    go func() {
        time.Sleep(dur)
        s.TaskQueue <- task
    }()
}

func main() {
    taskScheduler := NewTaskScheduler(1 * time.Second)
    taskScheduler.Start()

    for {
        task := PrintTask{Message: "Hey there"}
        taskScheduler.ScheduleOnce(0, task) // Schedules the task to be executed immediately
        time.Sleep(3 * time.Second) // Wait for 3 seconds before scheduling the next task
    }
}

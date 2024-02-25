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

func main() {
	fmt.Print("Task name: ")
	var task string
	fmt.Scanln(&task)
	fmt.Printf("Task has name: %s.\n", task)

	currentTime := time.Now()
	fmt.Printf("Current time: %s\n", currentTime.Format("2006–01–02 15:04:05"))
}

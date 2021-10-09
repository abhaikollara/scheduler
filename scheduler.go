package scheduler

import (
	"net/http"
	"time"
)

type Scheduler interface {
	Start() error
	Schedule(time.Time, *http.Request) error
}

type SimpleScheduler struct {
	taskStore TaskStore
}

func (s *SimpleScheduler) Schedule(task Task) error {
	return s.taskStore.SaveTask(task)

}

func (s *SimpleScheduler) Start() error {
	var ticker = time.NewTicker(time.Second)
	for {
		t := <-ticker.C
		go func() {
			taskList, _ := s.taskStore.GetTasksFor(t)

			for _, task := range taskList {
				go task.Execute()
			}
		}()

	}
}

func NewSimpleScheduler(taskStore TaskStore) SimpleScheduler {
	return SimpleScheduler{taskStore: taskStore}
}

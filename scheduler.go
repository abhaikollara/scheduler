package scheduler

import (
	"time"
)

type Scheduler interface {
	Start() error
	Schedule(time.Time, Task) error
}

type SimpleScheduler struct {
	taskStore TaskStore
}

func (s *SimpleScheduler) Schedule(tasks ...Task) error {
	var err error
	for _, t := range tasks {
		err = s.taskStore.SaveTask(t)
		if err != nil {
			return err
		}
	}

	return nil

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

func NewSimpleScheduler() SimpleScheduler {
	return SimpleScheduler{taskStore: NewInMemoryStore()}
}

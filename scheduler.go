package main

import (
	"net/http"
	"time"
)

type Scheduler interface {
	Start() error
	Schedule(time.Time, *http.Request) error
}

type VanillaScheduler struct {
	taskStore TaskStore
}

func (s *VanillaScheduler) Schedule(t time.Time, req *http.Request) error {
	var err error

	task := HTTPRequestTask{ScheduledAt: t, Request: req}
	s.taskStore.SaveTask(task)

	return err
}

func (s *VanillaScheduler) Start() error {
	var ticker = time.NewTicker(time.Second)
	for {
		t := <-ticker.C
		taskList, _ := s.taskStore.GetTasksFor(t)

		for _, task := range taskList {
			go task.Execute()
		}
	}
}

func NewVanillaScheduler(taskStore *TaskStore) (VanillaScheduler, error) {
	var s VanillaScheduler

	s.taskStore = *taskStore

	return s, nil
}

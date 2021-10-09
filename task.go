package main

import (
	"net/http"
	"time"
)

type Task interface {
	GetScheduleTime() time.Time
	Execute() error
}

type HTTPRequestTask struct {
	ScheduledAt time.Time
	Request     *http.Request
}

func (t HTTPRequestTask) GetScheduleTime() time.Time {
	return t.ScheduledAt
}

func (t HTTPRequestTask) Execute() error {
	client := http.Client{}

	_, err := client.Do(t.Request)
	if err != nil {
		return err
	}

	return nil
}

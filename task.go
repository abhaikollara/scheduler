package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task interface {
	GetScheduleTime() time.Time
	Execute() error
}

type PrintTask struct {
	ScheduledAt time.Time
	Message     string
}

func NewPrintTask(t time.Time, message string) PrintTask {
	return PrintTask{ScheduledAt: t, Message: message}
}

func (t PrintTask) GetScheduleTime() time.Time {
	return t.ScheduledAt
}

func (t PrintTask) Execute() error {
	fmt.Println(t.Message)
	return nil
}

type HTTPRequestTask struct {
	ScheduledAt time.Time
	Request     *http.Request
}

func NewHTTPRequestTask(t time.Time, req *http.Request) HTTPRequestTask {
	return HTTPRequestTask{ScheduledAt: t, Request: req}
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

func NewHTTPPostTask(t time.Time, payload json.RawMessage, url string) HTTPRequestTask {
	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))
	return NewHTTPRequestTask(t, req)
}

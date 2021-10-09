package scheduler

import "time"

type TaskStore interface {
	SaveTask(Task) error
	GetTasksFor(time.Time) ([]Task, error)
}

type InMemoryStore struct {
	store map[int64][]Task
}

func (s *InMemoryStore) SaveTask(t Task) error {
	var err error

	key := t.GetScheduleTime().Unix()

	if s.store[key] != nil {
		s.store[key] = append(s.store[key], t)
	} else {
		s.store[key] = []Task{t}
	}

	return err
}

func (s *InMemoryStore) GetTasksFor(t time.Time) ([]Task, error) {
	var scheduledTasks []Task

	if tasksList, ok := s.store[t.Unix()]; ok {
		scheduledTasks = tasksList
	} else {
		scheduledTasks = []Task{}
	}

	return scheduledTasks, nil
}

func NewInMemoryStore() *InMemoryStore {
	ms := InMemoryStore{}
	ms.store = make(map[int64][]Task)

	return &ms
}

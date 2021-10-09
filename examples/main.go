package main

import (
	"time"

	"github.com/abhaikollara/scheduler"
)

func main() {
	ts := scheduler.NewInMemoryStore()
	s := scheduler.NewSimpleScheduler(ts)

	t1 := scheduler.NewPrintTask(time.Now().Add(time.Second*1), "lorem")
	t2 := scheduler.NewPrintTask(time.Now().Add(time.Second*2), "ipsum")
	t3 := scheduler.NewPrintTask(time.Now().Add(time.Second*3), "dolores")

	s.Schedule(t1)
	s.Schedule(t2)
	s.Schedule(t3)

	s.Start()
}

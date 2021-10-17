package main

import (
	"time"

	"github.com/abhaikollara/scheduler"
)

func main() {
	s := scheduler.NewSimpleScheduler()

	t1 := scheduler.NewPrintTask(time.Now().Add(time.Second*1), "lorem")
	t2 := scheduler.NewPrintTask(time.Now().Add(time.Second*3), "ipsum")
	t3 := scheduler.NewPrintTask(time.Now().Add(time.Second*2), "dolores")

	s.Schedule(t1, t2, t3)

	s.Start()
}

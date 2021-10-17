# Scheduler
A simple task scheduler library written in Go

## Example
```go
package main

import (
	"time"

	"github.com/abhaikollara/scheduler"
)

func main() {
	s := scheduler.NewSimpleScheduler()

	t1 := scheduler.NewPrintTask(time.Now().Add(time.Second*1), "lorem")
	t2 := scheduler.NewPrintTask(time.Now().Add(time.Second*2), "ipsum")
	t3 := scheduler.NewPrintTask(time.Now().Add(time.Second*3), "dolores")

	s.Schedule(t1, t2, t3)

	s.Start()
}

}
```

### Output
```
lorem: 2021-10-10 01:10:34.707172 +0530 IST m=+1.000395209
ipsum: 2021-10-10 01:10:35.707173 +0530 IST m=+2.000396667
dolores: 2021-10-10 01:10:36.707174 +0530 IST m=+3.000396834
```

package main

import (
	"net/http"
	"time"
)

func main() {
	s, _ := NewVanillaScheduler()

	req, _ := http.NewRequest("GET", "http://example.org/", nil)

	s.Schedule(time.Now().Add(time.Second*5), req)

	s.Start()
}

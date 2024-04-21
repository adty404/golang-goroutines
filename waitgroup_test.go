package belajar_golang_goroutines

import (
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	println("Run Asynchronous")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go RunAsynchronous(group)
	}

	group.Wait()
	println("Finish")
}
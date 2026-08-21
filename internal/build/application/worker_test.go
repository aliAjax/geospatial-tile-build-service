package application

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPoolStopWaitsForAcceptedTasks(t *testing.T) {
	pool := NewWorkerPool(3, 16)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool.Start(ctx)
	var done atomic.Int32
	start := make(chan struct{})
	var submitted sync.WaitGroup
	for producer := 0; producer < 2; producer++ {
		submitted.Add(1)
		go func() {
			defer submitted.Done()
			<-start
			for i := 0; i < 6; i++ {
				for !pool.Submit(Task{ID: "task", Run: func(context.Context) error { time.Sleep(2 * time.Millisecond); done.Add(1); return nil }}) {
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}
	close(start)
	submitted.Wait()
	pool.Stop()
	if got := done.Load(); got != 12 {
		t.Fatalf("Stop returned before accepted tasks finished: %d", got)
	}
}

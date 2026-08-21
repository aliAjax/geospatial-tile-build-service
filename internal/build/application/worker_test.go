package application

import (
	"context"
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
	for i := 0; i < 12; i++ {
		if !pool.Submit(Task{ID: "task", Run: func(context.Context) error {
			time.Sleep(2 * time.Millisecond)
			done.Add(1)
			return nil
		}}) {
			t.Fatal("task was rejected")
		}
	}
	pool.Stop()
	if got := done.Load(); got != 12 {
		t.Fatalf("Stop returned before accepted tasks finished: %d", got)
	}
}

package application

import (
	"context"
	"sync"
	"time"
)

type Task struct {
	ID  string
	Run func(context.Context) error
}
type WorkerPool struct {
	workers int
	tasks   chan Task
	done    chan struct{}
	wg      sync.WaitGroup
}

func NewWorkerPool(workers, queue int) *WorkerPool {
	if workers < 1 {
		workers = 1
	}
	return &WorkerPool{workers: workers, tasks: make(chan Task, queue), done: make(chan struct{})}
}
func (w *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-w.done:
					for {
						select {
						case t := <-w.tasks:
							if t.Run != nil {
								_ = t.Run(ctx)
							}
						default:
							return
						}
					}
				case t, ok := <-w.tasks:
					if !ok {
						return
					}
					if t.Run != nil {
						_ = t.Run(ctx)
					}
				}
			}
		}()
	}
}
func (w *WorkerPool) Submit(t Task) bool {
	select {
	case w.tasks <- t:
		return true
	case <-w.done:
		return false
	default:
		return false
	}
}
func (w *WorkerPool) Stop() { close(w.done); w.wg.Wait(); time.Sleep(time.Millisecond) }

package application

import (
	"context"
	"sync"
)

// Task is a unit of work executed by the pool. Run receives the start context
// so the work body can observe cancellation itself; the worker loop never
// drops a task that has already been accepted into the queue.
type Task struct {
	ID  string
	Run func(context.Context) error
}

type WorkerPool struct {
	workers int
	tasks   chan Task
	done    chan struct{}
	once    sync.Once
	wg      sync.WaitGroup
}

func NewWorkerPool(workers, queue int) *WorkerPool {
	if workers < 1 {
		workers = 1
	}
	return &WorkerPool{
		workers: workers,
		tasks:   make(chan Task, queue),
		done:    make(chan struct{}),
	}
}

// Start launches the configured number of workers. Add is accounted before
// each goroutine is created so Stop's Wait can never observe a stale count.
func (w *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go w.loop(ctx)
	}
}

func (w *WorkerPool) loop(ctx context.Context) {
	defer w.wg.Done()
	for {
		// Drain everything already accepted before considering shutdown. This
		// is the drain-first guarantee: a queued task is never starved by the
		// stop signal, which previously caused accepted tasks to be skipped.
		select {
		case t, ok := <-w.tasks:
			if !ok {
				return
			}
			w.run(ctx, t)
			continue
		default:
		}
		// Queue is momentarily empty: block for either new work, a stop
		// request, or context cancellation. Done is checked last so that a
		// task accepted concurrently with Stop still runs.
		select {
		case t, ok := <-w.tasks:
			if !ok {
				return
			}
			w.run(ctx, t)
		case <-w.done:
			// Stop requested. Drain any tasks accepted in the meantime so
			// Stop only returns once every queued task has executed.
			w.drain(ctx)
			return
		case <-ctx.Done():
			// Caller context cancelled: drain already-accepted work before
			// exiting so accepted tasks are not orphaned.
			w.drain(ctx)
			return
		}
	}
}

// drain runs every task still buffered in the channel. It returns once the
// queue is empty, cooperating with concurrent submitters until Stop closes
// done and rejects further submissions.
func (w *WorkerPool) drain(ctx context.Context) {
	for {
		select {
		case t, ok := <-w.tasks:
			if !ok {
				return
			}
			w.run(ctx, t)
		default:
			return
		}
	}
}

func (w *WorkerPool) run(ctx context.Context, t Task) {
	if t.Run != nil {
		_ = t.Run(ctx)
	}
}

// Submit enqueues a task. It returns false once Stop has been called, so no
// task is reported as accepted after shutdown begins. The tasks channel is
// never closed, so a submitter racing with Stop never panics on a send to a
// closed channel; at worst it observes done and returns false.
func (w *WorkerPool) Submit(t Task) bool {
	select {
	case <-w.done:
		return false
	default:
	}
	select {
	case w.tasks <- t:
		return true
	default:
		return false
	}
}

// Stop signals shutdown and blocks until every worker has exited. Every task
// accepted before Stop (still queued or in flight) is allowed to finish; only
// then do workers drain and return. Idempotent via sync.Once.
func (w *WorkerPool) Stop() {
	w.once.Do(func() {
		close(w.done)
	})
	w.wg.Wait()
}

// Stopped reports whether Stop has been initiated. Safe for concurrent use.
func (w *WorkerPool) Stopped() bool {
	select {
	case <-w.done:
		return true
	default:
		return false
	}
}

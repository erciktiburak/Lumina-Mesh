package worker

import (
	"context"
	"log"
	"sync"
)

type Task struct {
	ID      string
	Payload []byte
}

type Pool struct {
	tasks       chan Task
	workerCount int
	wg          sync.WaitGroup
}

func NewPool(workerCount int, bufferSize int) *Pool {
	return &Pool{
		tasks:       make(chan Task, bufferSize),
		workerCount: workerCount,
	}
}

func (p *Pool) Start(ctx context.Context) {
	for i := 0; i < p.workerCount; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}
	log.Printf("Worker pool started with %d workers", p.workerCount)
}

func (p *Pool) Stop() {
	close(p.tasks)
	p.wg.Wait()
	log.Println("Worker pool stopped")
}

func (p *Pool) Submit(t Task) {
	p.tasks <- t
}

func (p *Pool) worker(ctx context.Context, id int) {
	defer p.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			p.process(id, task)
		}
	}
}

func (p *Pool) process(workerID int, t Task) {
	log.Printf("[Worker-%d] Processing task %s", workerID, t.ID)
	// Simulate processing time
}

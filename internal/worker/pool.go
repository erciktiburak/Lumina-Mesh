package worker

import (
	"context"
	"log"
	"sync"

	"github.com/erciktiburak/Lumina-Mesh/internal/messaging"
)

type Task struct {
	ID         string
	Payload    []byte
	RetryCount int
}

const (
	MaxRetries = 3
	DLQSubject = "mesh.dlq"
)

type Pool struct {
	tasks       chan Task
	workerCount int
	wg          sync.WaitGroup
	Processor   func(workerID int, t Task) error
	client      *messaging.Client
}

func NewPool(workerCount int, bufferSize int, client *messaging.Client) *Pool {
	p := &Pool{
		tasks:       make(chan Task, bufferSize),
		workerCount: workerCount,
		client:      client,
	}
	p.Processor = p.defaultProcess
	return p
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
			if err := p.Processor(id, task); err != nil {
				log.Printf("[Worker-%d] Task %s failed: %v", id, task.ID, err)
				if task.RetryCount < MaxRetries {
					task.RetryCount++
					log.Printf("[Worker-%d] Retrying task %s (%d/%d)", id, task.ID, task.RetryCount, MaxRetries)
					p.Submit(task)
				} else {
					log.Printf("[Worker-%d] Task %s exceeded max retries, sending to DLQ", id, task.ID)
					if p.client != nil {
						_ = p.client.Publish(DLQSubject, task.Payload)
					}
				}
			}
		}
	}
}

func (p *Pool) defaultProcess(workerID int, t Task) error {
	log.Printf("[Worker-%d] Processing task %s", workerID, t.ID)
	return nil
}

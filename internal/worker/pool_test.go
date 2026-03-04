package worker

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewPool(3, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	processedCount := int32(0)

	// Mock processing
	pool.Processor = func(workerID int, task Task) {
		atomic.AddInt32(&processedCount, 1)
	}

	pool.Start(ctx)

	for i := 0; i < 5; i++ {
		pool.Submit(Task{ID: "test", Payload: []byte("payload")})
	}

	time.Sleep(100 * time.Millisecond)
	pool.Stop()

	if atomic.LoadInt32(&processedCount) != 5 {
		t.Errorf("Expected 5 processed tasks, got %d", processedCount)
	}
}

package engine

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/erciktiburak/Lumina-Mesh/internal/runtime"
)

type ExecutionContext struct {
	mu      sync.RWMutex
	Outputs map[string]map[string]interface{}
}

type Executor struct {
	wasmRuntime *runtime.WasmRuntime
	DryRun      bool
}

func NewExecutor() *Executor {
	return &Executor{
		wasmRuntime: runtime.NewWasmRuntime(),
		DryRun:      false,
	}
}

func (e *Executor) Run(ctx context.Context, workflow *Workflow) error {
	if e.DryRun {
		log.Printf("[Executor] Dry Run mode enabled. Validating workflow %s...", workflow.Name)
	}
	execCtx := &ExecutionContext{
		Outputs: make(map[string]map[string]interface{}),
	}

	for _, step := range workflow.Steps {
		log.Printf("[Executor] Running step: %s", step.Name)

		// 1. Resolve inputs (from previous steps)
		// For now, let's just pass simple inputs.
		// In a full version, we'd parse ${steps.X.outputs.Y} strings.

		// 2. Execute step
		output, err := e.executeStep(ctx, step, execCtx)
		if err != nil {
			return fmt.Errorf("step %s failed: %w", step.Name, err)
		}

		// 3. Save outputs
		execCtx.mu.Lock()
		execCtx.Outputs[step.Name] = output
		execCtx.mu.Unlock()
	}

	log.Printf("[Executor] Workflow %s completed successfully", workflow.Name)
	return nil
}

func (e *Executor) executeStep(ctx context.Context, step Step, execCtx *ExecutionContext) (map[string]interface{}, error) {
	if e.DryRun {
		log.Printf("[Executor] [DryRun] Validating step: %s", step.Name)
		return map[string]interface{}{"status": "validated"}, nil
	}
	// Simple simulation of step execution
	// In reality, we'd check action type: wasm://, ai://, etc.
	return e.wasmRuntime.Execute(step.Action, step.Inputs)
}

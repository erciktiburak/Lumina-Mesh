package engine

import (
	"os"
	"testing"
)

func TestParseWorkflow(t *testing.T) {
	yamlContent := `
name: test-workflow
version: 1.0.0
steps:
  - name: step-1
    action: wasm://test.wasm
    inputs:
      key: value
`
	tmpfile, err := os.CreateTemp("", "workflow-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(yamlContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	workflow, err := ParseWorkflow(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseWorkflow failed: %v", err)
	}

	if workflow.Name != "test-workflow" {
		t.Errorf("Expected name 'test-workflow', got '%s'", workflow.Name)
	}

	if len(workflow.Steps) != 1 {
		t.Errorf("Expected 1 step, got %d", len(workflow.Steps))
	}

	if workflow.Steps[0].Name != "step-1" {
		t.Errorf("Expected step name 'step-1', got '%s'", workflow.Steps[0].Name)
	}
}

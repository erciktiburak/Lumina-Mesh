package engine

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Step struct {
	Name      string                 `yaml:"name"`
	Action    string                 `yaml:"action"`
	Condition string                 `yaml:"condition,omitempty"`
	Inputs    map[string]interface{} `yaml:"inputs,omitempty"`
	Outputs   map[string]string      `yaml:"outputs,omitempty"`
	Retry     int                    `yaml:"retry,omitempty"`
}

type Workflow struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Steps   []Step `yaml:"steps"`
}

func ParseWorkflow(path string) (*Workflow, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflow file: %w", err)
	}

	var workflow Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, fmt.Errorf("failed to parse workflow YAML: %w", err)
	}

	return &workflow, nil
}

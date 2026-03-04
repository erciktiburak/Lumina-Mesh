package api

import (
	"context"
	"log"

	"github.com/erciktiburak/Lumina-Mesh/pkg/api"
)

type Server struct {
	api.UnimplementedWorkflowServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) TriggerWorkflow(ctx context.Context, req *api.TriggerRequest) (*api.TriggerResponse, error) {
	log.Printf("[API] Triggering workflow: %s", req.WorkflowId)
	return &api.TriggerResponse{
		ExecutionId: "exec-123",
		Status:      "queued",
	}, nil
}

func (s *Server) GetWorkflowStatus(ctx context.Context, req *api.StatusRequest) (*api.StatusResponse, error) {
	log.Printf("[API] Getting status for execution: %s", req.ExecutionId)
	return &api.StatusResponse{
		Status: "completed",
	}, nil
}

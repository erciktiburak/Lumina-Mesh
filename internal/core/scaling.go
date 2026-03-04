package core

import (
	"log"
	"time"
)

type ScalingService struct {
	nodeCount int
}

func NewScalingService() *ScalingService {
	return &ScalingService{nodeCount: 1}
}

func (s *ScalingService) Monitor(load float64) {
	if load > 0.8 {
		s.scaleUp()
	} else if load < 0.2 {
		s.scaleDown()
	}
}

func (s *ScalingService) scaleUp() {
	s.nodeCount++
	log.Printf("[Scaling] High load detected! Scaling up to %d nodes...", s.nodeCount)
}

func (s *ScalingService) scaleDown() {
	if s.nodeCount > 1 {
		s.nodeCount--
		log.Printf("[Scaling] Low load detected! Scaling down to %d nodes...", s.nodeCount)
	}
}

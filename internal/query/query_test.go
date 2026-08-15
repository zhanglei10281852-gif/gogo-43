package query

import (
	"strings"
	"testing"
	"time"

	"QueueForge/internal/model"
)

func TestRouteReturnsErrorForNilJob(t *testing.T) {
	worker := model.Worker{ID: "worker", Queues: []string{"default"}, Capacity: model.Resources{Slots: 1}}
	job := &model.Job{ID: "job", Queue: "default", State: model.StateReady, AvailableAt: time.Unix(100, 0).UTC()}
	_, err := Route([]*model.Job{job, nil}, worker, time.Unix(100, 0).UTC())
	if err == nil || !strings.Contains(err.Error(), "nil job") {
		t.Fatalf("Route() error = %v, want nil-job error", err)
	}
}

package worker_pool

import (
	"testing"
	"time"
)

type MockTask struct {
	ID         int
	Executed   bool
	Failed     bool
	FailureErr error
}

func (t *MockTask) Execute() error {
	t.Executed = true
	return t.FailureErr
}

func (t *MockTask) OnFailure(err error) {
	t.Failed = true
	t.FailureErr = err
}

func TestWorkerPool(t *testing.T) {
	var tasks []*MockTask
	numTask := 20
	for i := 0; i < numTask; i++ {
		tasks = append(tasks, &MockTask{ID: i})
	}

	p, err := NewSimplePool(5, numTask)
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()

	for _, j := range tasks {
		p.AddWork(j)
	}
	// Sleep for a short duration to allow tasks to be processed
	time.Sleep(500 * time.Millisecond)

	// Check if all tasks were executed
	for i, task := range tasks {
		if !task.Executed {
			t.Errorf("Task %d was not executed", i)
		}
	}
	p.Stop()
}

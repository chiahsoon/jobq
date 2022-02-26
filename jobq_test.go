package jobq_test

import (
	"testing"
	"time"

	"github.com/chiahsoon/jobq"
	"github.com/stretchr/testify/assert"
)

type SetNumTask struct {
	Name    string
	SrcVal  int
	DestVal *int
}

func (t *SetNumTask) Desc() string { return t.Name }
func (t *SetNumTask) Run()         { (*t.DestVal) = t.SrcVal }

func TestBasicSetNumTask(t *testing.T) {
	assert := assert.New(t)

	originalVal := 10
	val := originalVal
	task := SetNumTask{DestVal: &val, SrcVal: originalVal + 1}
	job := jobq.Job{
		Label:    "Basic Job",
		Tasks:    []jobq.Task{&task},
		Priority: 1,
	}
	queue := jobq.PriorityQueue{}
	jq := jobq.NewJobQ(0, &queue)

	jq.Watch(false)
	jq.Add(&job)
	time.Sleep(time.Millisecond) // Wait for job to end

	assert.Equal(originalVal+1, val, "Increment task should be sucessful")
}

func TestPriority(t *testing.T) {
	assert := assert.New(t)

	val := 10
	higherPriorityTask := SetNumTask{SrcVal: 20, DestVal: &val}
	higherPriorityJob := jobq.Job{
		Label:    "Lower Priority Job",
		Tasks:    []jobq.Task{&higherPriorityTask},
		Priority: 2,
	}
	lowerPriorityTask := SetNumTask{SrcVal: 15, DestVal: &val}
	lowerPriorityJob := jobq.Job{
		Label:    "Higher Priority Job",
		Tasks:    []jobq.Task{&lowerPriorityTask},
		Priority: 1,
	}

	queue := jobq.PriorityQueue{}
	queue.PushJob(&lowerPriorityJob)
	queue.PushJob(&higherPriorityJob)
	jq := jobq.NewJobQ(0, &queue)

	jq.Watch(false)
	jq.Add(&higherPriorityJob)
	time.Sleep(time.Millisecond)

	assert.Equal(higherPriorityTask.SrcVal, val, "Higher priority task should have ran first")
	assert.NotEqual(lowerPriorityTask.SrcVal, val, "Lower priority task should not have ran first")
}

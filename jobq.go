package jobq

import (
	"time"
)

type JobQ struct {
	// Interval to waitInterval before checking if there are jobs
	waitInterval time.Duration
	queue        Queue
}

func NewJobQ(waitInterval time.Duration, queue Queue) *JobQ {
	return &JobQ{waitInterval: waitInterval, queue: queue}
}

func (jq *JobQ) Add(job *Job) {
	jq.queue.PushJob(job)
}

func (jq *JobQ) Remove() *Job {
	return jq.queue.PopJob()
}

func (jq *JobQ) Watch() {
	// Stop goroutine ref: https://stackoverflow.com/a/37997989/7550732
	for {
		// Run all ready jobs
		for jq.queue.HasReadyJob() {
			jq.queue.PopJob().Run()
		}
		time.Sleep(jq.waitInterval)
	}
}

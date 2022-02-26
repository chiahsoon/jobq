package jobq

import (
	"time"
)

type JobQ struct {
	pollInterval time.Duration
	queue        Queue
}

func NewJobQ(pollInterval time.Duration, queue Queue) *JobQ {
	return &JobQ{pollInterval: pollInterval, queue: queue}
}

func (jq *JobQ) Add(job *Job) {
	jq.queue.PushJob(job)
}

func (jq *JobQ) Remove() *Job {
	return jq.queue.PopJob()
}

func (jq *JobQ) Watch(block bool) {
	// Stop goroutine ref: https://stackoverflow.com/a/37997989/7550732
	watch := func(jq *JobQ) {
		for {
			if jq.queue.HasReadyJob() {
				job := jq.queue.PopJob()
				_ = job.Run()
			}
			time.Sleep(jq.pollInterval)
		}
	}

	if block {
		watch(jq)
		return
	}
	go watch(jq)
}

package jobq

type Queue interface {
	PushJob(job *Job)
	PopJob() *Job
	UpdateJob(item *Job, label string, tasks []Task, priority int)
	Len() int
	Peek() *Job
	HasReadyJob() bool
}

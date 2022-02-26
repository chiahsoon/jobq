package jobq

type Job struct {
	Label    string
	Tasks    []Task
	Priority int
	index    int
}

func (j *Job) Run() {
	for _, task := range j.Tasks {
		task.Run()
	}
}

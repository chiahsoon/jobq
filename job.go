package jobq

type Job struct {
	Label    string
	Tasks    []Task
	Priority int
	index    int
}

func (j *Job) Run() error {
	for _, task := range j.Tasks {
		if err := task.Run(); err != nil {
			return err
		}
	}
	return nil
}

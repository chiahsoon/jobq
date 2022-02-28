package jobq

import "fmt"

type Job struct {
	Label    string
	Tasks    []Task
	Priority int
	index    int
}

func (j *Job) Run() {
	fmt.Println("Running Job: ", j.Label)
	for _, task := range j.Tasks {
		fmt.Println("Running Task: ", task.Desc())
		task.Run()
	}
}

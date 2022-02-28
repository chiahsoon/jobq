package jobq

import (
	"fmt"
	"sync"
)

type Job struct {
	Label    string
	Tasks    []Task
	Priority int
	index    int
}

func (j *Job) Run() {
	fmt.Println("Running Job: ", j.Label)
	if len(j.Tasks) == 0 {
		fmt.Println("There are no tasks to run.")
		return
	}
	fmt.Printf("Running a total of %d tasks\n", len(j.Tasks))

	wg := sync.WaitGroup{}
	for idx, task := range j.Tasks {
		wg.Add(1)
		go func(taskIdx int, task Task) {
			fmt.Printf("Running Task %d: %s\n", taskIdx, task.Desc())
			task.Run()
			wg.Done()
		}(idx, task)
	}
	wg.Wait()
}

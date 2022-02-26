package jobq

import (
	"container/heap"
)

type PriorityQueue []*Job

// Queue Methods
func (pq *PriorityQueue) PushJob(job *Job) {
	heap.Push(pq, job)
}

func (pq *PriorityQueue) PopJob() *Job {
	return heap.Pop(pq).(*Job)
}

func (pq *PriorityQueue) UpdateJob(item *Job, label string, tasks []Task, priority int) {
	item.Label = label
	item.Tasks = tasks
	item.Priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Peek() *Job {
	if pq.Len() <= 0 {
		return nil
	}

	return (*pq)[0]
}

func (pq *PriorityQueue) HasReadyJob() bool {
	return pq.Len() > 0
}

// For both Heap & Queue implementation
func (pq PriorityQueue) Len() int { return len(pq) }

// Heap Methods
func (pq PriorityQueue) Less(i, j int) bool {
	// container/heap is a minheap by default;
	// return True when 'j' should be popped first
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Job)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

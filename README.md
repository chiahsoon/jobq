# JobQ

> *Queue == Que == Q*

A basic job queue implemented in Go.

### Idea
* `Task` is the basic unit of work that we want to run.
* `Job` contains >= 1 related `Task`(s).
* `Queue` contains jobs and determines how/when to execute certain `Job`s.
* `JobQ` is just a runner that runs ready `Job`s.

### Current Features
1. Both `Queue` and `Task` are interfaces that you can easily implement and provide additional fields to. 
2. `PriorityQueue` is provided as a default to use with `JobQ`. 
    * It is heap-based, so it implements the methods necessary for [container/heap](https://pkg.go.dev/container/heap).
3. Priority can be anything as long as they can be converted/encoded to `int` (string, date, numbers, etc.)

### Examples
Do take a look at `jobq_test.go` to see how `JobQ` is used.
# JobQ

> *Queue == Que == Q*

A basic job queue implemented in Go.

#### Idea
* `Task` is the basic unit of work that we want to run.
* `Job` contains >= 1 related `Task`(s).
* `Queue` contains jobs and determines how/when to execute certain `Job`s.
* `JobQ` is just a runner that runs ready `Job`s.

#### Current Features
1. Both `Queue` and `Task` are interfaces that you can easily implement. 
2. Priority can be anything - string, date, number; you just need to convert them to `int`.
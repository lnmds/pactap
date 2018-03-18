package main

/*
    Job management utilities.
*/

type JobState int

const (
    SCHEDULED = iota
    RUNNING
    DONE
    FAILURE
)

type Job struct {
    name string

    description string

    state JobState
}

type JobManager struct {
    name string

    jobs []Job
}

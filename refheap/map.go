package main

import "fmt"

type Job struct {
	Name string

	Priority int
}

var jobs = make(map[string]Job)

func main() {
	fmt.Println("Starting...")
	jobs["job1"] = Job{Name: "First Job"}
	p()

	job := jobs["job1"]
	job.Priority = 10
	p()

	// Workaround for issue 3117
	jobs["job1"] = job
	p()

	jobs["job1"].(*Priority) = 15
	p()
}

func p() {
	fmt.Printf("Jobs: %s\n", jobs)
}

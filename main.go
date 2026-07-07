package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID        int
	Name      string
	Timestamp time.Time
	Payload   any
	Status    string
	Attempts  int
}

func main() {
	job := Job{ID: 1, Name: "Alice", Timestamp: time.Now(), Payload: 42, Status: "Pending", Attempts: 1}

	result := work(job)
	if result == "Pending" {
		fmt.Println("Job is pending. Try again after some time")
	} else {
		fmt.Println("Job is completed after ", job.Attempts, " attempts")
	}

}

func work(job Job) string {
	time.Sleep(2 * time.Second)
	job.Status = "Completed"
	return job.Status
}

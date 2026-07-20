package main

import (
	"fmt"
	"sync"
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
	job_ch := make(chan *Job)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go work(i, job_ch, &wg)
	}

	jobs := []*Job{
		{ID: 1, Name: "resize-image", Timestamp: time.Now(), Payload: "img1.png", Status: "Pending", Attempts: 0},
		{ID: 2, Name: "send-email", Timestamp: time.Now(), Payload: "user@example.com", Status: "Pending", Attempts: 0},
		{ID: 3, Name: "generate-report", Timestamp: time.Now(), Payload: "Q3-report", Status: "Pending", Attempts: 0},
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main about to send jobs now")
	for _, j := range jobs {
		job_ch <- j
	}
	close(job_ch)

	wg.Wait()
	fmt.Println("main: all workers done")
}

func work(id int, job_ch <-chan *Job, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id, "waiting for a job")

	for j := range job_ch {
		fmt.Printf("Worker %d working on job %d, name %d\n", id, j.ID, j.Name)
		time.Sleep(2 * time.Millisecond)
		fmt.Printf("worker %d finished job %d\n", id, j.ID)
	}

	fmt.Printf("worker %d: channel closed, exiting \n", id)

}

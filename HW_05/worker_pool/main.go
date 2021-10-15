package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

const (
	numOfJobs    = 7
	numOfWorkers = 4
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker with id = %v started job = %v\n", id, j)
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		fmt.Printf("Worker with id = %v finished job = %v\n", id, j)
		results <- j
	}
	wg.Done()
}

func main() {
	wg.Add(numOfWorkers)
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()

	// Also, we can use this to print results
	//for i := 1; i <= numJobs; i++ {
	//	<-results
	//}
}

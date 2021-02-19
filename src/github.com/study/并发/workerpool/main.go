package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func producer(j chan<- *job) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Int63()
		newJob := &job{
			value: random,
		}
		j <- newJob
		time.Sleep(time.Second * 3)
	}
}

func consumer(ch1 <-chan *job, ch2 chan<- *result) {
	defer wg.Done()
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: *job,
			sum: sum,
		}
		ch2 <- newResult
	}
}

func main() {
	wg.Add(1)
	go producer(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go consumer(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d,sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}

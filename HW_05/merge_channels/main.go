package main

import (
	"fmt"
	"sync"
)

func generateChannel(start, stop int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			ch <- i
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

func merge(cs ...<-chan int) <-chan int {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	for _, c := range cs {
		wg.Add(1)
		localCh := c
		go func() {
			defer wg.Done()
			for v := range localCh {
				ch <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	c1 := generateChannel(1, 6)
	c2 := generateChannel(6, 11)
	for v := range merge(c1, c2) {
		fmt.Printf("%v ", v)
	}
}

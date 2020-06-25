package main

import (
	"strconv"
	"sync"
)

type Counter struct {
	Value int
	sync.Mutex
}

func main() {
	waiter := &sync.WaitGroup{}
	counter := Counter{}
	for index:=0; index < 20000; index++ {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			counter.Lock()
			counter.Value++
			counter.Unlock()
		}()
	}

	waiter.Wait()
	println("Counter is: " + strconv.Itoa(counter.Value))
}

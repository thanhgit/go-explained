package main

import (
	"fmt"
	"strconv"
	"sync"
)

func output(channel chan int) {
	for value := range channel {
		fmt.Println("Value: " + strconv.Itoa(value))
	}
}

func main() {
	waiter := &sync.WaitGroup{}
	var channel = make(chan int)
	go output(channel)

	for index := 0; index < 100; index++ {
		waiter.Add(1)
		go func(index int) {
			defer waiter.Done()
			channel <- index
		}(index)
	}

	waiter.Wait()
}

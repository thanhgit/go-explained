package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numberOfRequest := 1
	sem := make(chan int, numberOfRequest)
	//done := time.After(10*time.Second)
	var requests [5]int
	for index := 0; index < 5; index++ {
		requests[index] = rand.Int()+100
	}

	for _,req := range requests {
		sem <- 1
		go func(req int) {
			if len(sem) == numberOfRequest {
				fmt.Println("Full")
			}

			println(req)
			<-sem
		}(req)
	}

	//time.Sleep(time.Second)
}

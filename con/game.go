package main

import (
	"fmt"
	"os"
	"time"
)

func print(channel chan string) {
	for value:= range channel {
		fmt.Println(value)
	}
}


func main()  {
	done := time.After(5*time.Second)
	channel := make(chan string)
	defer close(channel)
	go func() {
		channel <- "Hello"
		channel <- "Hi"
		channel <- "halhala"
	}()

	for {
		select {
		case messsage := <-channel:
			fmt.Println(messsage)
			case <-done:
				fmt.Println("Done")
				os.Exit(0)
		}
	}

	//runtime.Gosched()
}

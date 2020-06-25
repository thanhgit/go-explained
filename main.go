package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func input(in chan int) {
	for index := 0; index < 10; index++ {
		go func(in chan int, index int) {
			in <- index
			time.Sleep(time.Millisecond * 10)
		}(in, index)

	}
}

func output(out chan int) {
	for value := range out {
		fmt.Println("hello " + strconv.Itoa(value))
	}

}

func main() {
	defer func() {
		if err:=recover(); err!= nil {
			fmt.Println("Panic is executed")
		}
	}()

	channel := make(chan int)
	go output(channel)

	//go input(channel)
	scanner := bufio.NewScanner(os.Stdin)
	options := 0
	fmt.Println("MENU")
	fmt.Print("Chose: 0 or !0: ")
	for options == 0 && scanner.Scan() {

		if scanner.Text() == "0" {
			options = 1
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				//fmt.Println("not number")
				//panic(errors.New("Not number"))
				panic(fmt.Errorf("Error: %s", "Not number"))
				continue
			}
			channel <- value
			fmt.Println("MENU")
			fmt.Print("Chose: 0 or !0: ")
		}
	}
}

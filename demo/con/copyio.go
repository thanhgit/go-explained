package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {
	waiter := &sync.WaitGroup{}
	waiter.Add(1)
	go echo(waiter,os.Stdin, os.Stdout)
	//time.Sleep(10*time.Second)
	//os.Exit(0)
	//runtime.Gosched()

	time.Sleep(time.Second)
	waiter.Wait()
}

func echo(waiter *sync.WaitGroup, in io.Reader, out io.Writer) {
	defer waiter.Done()
	//io.Copy(out, in)
	//scanner := bufio.NewScanner(in)
	//for scanner.Scan() {
	//	fmt.Println("Get: " + scanner.Text())
	//}

	for index:=0; index < 10; index++ {
		fmt.Println(index)
	}

}

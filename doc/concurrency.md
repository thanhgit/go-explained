## Conccurency in go

### Using runtime.Gosched()
- This function allow runtime an opportunity to execute other goroutines before it exits
- The other way is using time.Sleep()

### Using sync.WaitGroup
```text
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
```
### Using mutex for race condition
````text
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
````
### Using multi channel with select
```text
done := time.After(5*time.Second)
channel := make(chan string)
defer close(channel)
go func() {
    channel <- "Hello"
    channel <- "How are you?"
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
```
### Using buff channel to limit number of request or lock code space
```text
var sem = make(chan int, MaxNumRequest)

func Serve(queue chan *Request) {
        for req := range queue {
                sem <- 1
                go func(req *Request) {
                        process(req)
                        <-sem
                }(req)
    }
}
```
### Close channel with Close(ch)
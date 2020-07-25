# Go basic knowledge

## Libraries
* ### Against Cross Site Script with 
	https://golang.org/pkg/html/template/
	https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.3.html
## Tools
* ### Rewites code into the standard format
```text
gofmt -s -w main.go
```
* ### Manages the insertion and removal of import declarations as needed 
```text
goimports -w main.go
```
## Common function
- len(_var)
- cap(_var), 
- arr := make([]int, 5)
- sa := append(s, 0)
- for index, value := range { ... }
- _map := make(map[string][int])
    - map['helo'] = 1
    - map['hi'] = 1
    - hello = map['hello']
    - delete(map, 'hi')
- errors.New('--str---') with import "errors"
- err := fmt.Errorf("Error: %q", str)
* ### HOC
```text
func compute(x int, y int, add func(int, int) int) int {
    ...
    return func(x, y)
}

compute(1, 2, func(x int, y int) {
    return x + y
})
```
* ### Closure
```text
func adder() func(int) int {
    sum:=0;
    return func(x int) int {
        sum += x
        return sum
    }
}
```
* ### Non-struct type
```text
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
```
* ### Interface 
- An interface type is defined as a set of method signatures 
- A value of interface type can hold any value that implements those methods
- Interface values - (value, type) using fmt.Printf("(%v, %T)\n", i, i)
- interface{} - empty interface, is used by code that handles values of unknown type
* ### Type assertions
```text
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)

f = i.(float64) // panic
fmt.Println(f)
```
OR
```text
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```
* ### Convert with fmt.Sprintf
```text
str := fmt.Sprintf("%d.%d.%d.%d", int(ipAddr[0]), int(ipAddr[1]), int(ipAddr[2]), int(ipAddr[3]))

// [127 0 0 1] -> 127.0.0.1
``
* ### Stringer in detail
```text
String() string
Error() string 
```
* ### Channels
```text
ch <- v // send v to ch channel
value := <- ch
// create channel
ch := make(chan int)
// close channel
deffer close(ch)
// assert channel
value, ok := <-ch 
- ok == false -> value =

// using select 
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
// Using mutex 
- Lock
- UnLock
```

# Golang best practise

```text
const DEBUG = true

If DEBUG {
    defer func(v int) {
        fmt.Println(“receive:”, v)
    }(value)
}
```
## Custom error type
```text
type MyError struct{
    Message string
}
func (m MyError) Error() string {
    return m.Message
}
var err error = MyError{Message: “Hello my first custom error”}
```
## Panic and recover 
- panic is similar to `throw error`
- recover is similar to `catch error`

## Directory structure
```text
project/
    model/
    repository/
    handler/
    driver/
    main.go
```
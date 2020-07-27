# Go formatting

```go
type Student struct {
   Id           int64
   Name         string
   Grade        int
   CurrentScore float32
}

// create instance
student := Student{
   Id:           50,
   Name:         "John Smith",
   Grade:        5,
   CurrentScore: 3.8,
}
```
## Object Formatting
* ### Print the struct 
```go
result := fmt.Sprintf("%v", student)

// {50 John Smith 5 3.8}
```
* ### Print the struct with field names
```go
result := fmt.Sprintf("%+v", student)

// {Id:50 Name:John Smith Grade:5 CurrentScore:3.8}
```
* ### Print the type of the value
```go
result := fmt.Sprintf("%T", student)

// main.Student
```
* ### Print the type and actual value
```go
result := fmt.Sprintf("%#v")

// main.Student{Id:50, Name:"John Smith", Grade:5, CurrentScore:3.8}
```
## String Formatting
```go
sample := "I love Golang!"
```
* ### Print the string
```go
result := fmt.Sprintf("Sample is: %s", sample)

// Sample is : I love Golang!
```
* ### Print the string with quotes
```go
result := fmt.Sprintf("Sample is: %q", sample)

// Sample is : "I love Golang!"
```
## Boolean formatting
```go
sample := true
```
* ### Print boolean
```go
result := fmt.Sprintf("Sample is: %t", sample)

// Sample is : true
```
* ### Integer formatting
```text
# integer base-10
42 -> %d -> 42
# integer base-2
42 -> %b -> 101010
```
* ### Float formatting
```text
# Print float decimal format
3.1415926 -> %f -> 3.1415926
# Print float in decimal format with precision
3.1415926 -> %.9f -> 3.141592600
```
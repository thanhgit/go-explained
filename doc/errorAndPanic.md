## Handling error and panic
Go distinguishes between error and panic 
- Error indicate a particular task couldn't is completed successfully
- Panic indicate that a severe event occurred, such as error programmer

### Create error
- #### error.New("Error at main.go")
- #### fmt.Errorf(...)

### Best practise for handling error
* #### func convert(str string) (data, err) { ... }
* #### Custom error types
```text
type error interface {
    Error() string
}

type ParseError struc {
    Message string
    Line, Char int
}

func (p *ParseError)Error() string {
    format := "%s on Line %d, Char %d"
    return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
```
* #### Error variables
    * ##### Create package-scoped error variables, return a certain error occurs, such as: io.EOF, io.ErrNoProcess
    ```text
    var ErrorTimeout = errors.New("The request timeout")
    var ErrorReject = errors.New("The request was rejected")
    ```
### Panic system 
- Indicate that something has gone wrong that system can't continue to function
- When panic occurs, 
    - Looking for handlers for that panic. 
    - If not handler is found, unwinds to the top of the function stack and stop program
* #### Using panic(interface{})
    * ##### Best thing pass to panic function is a error
    * ```panic(errors.New("Error divide for 0"))```
    * ```panic(fmt.Errorf("Error: %s", "Divide for 0"))```
    * #### Recovering from panics
        ```text
        defer func(){
            if err:=recover(); err!=nill {
                fmt.Println("Error divide for 0")
            }
        }()
        ```
    * ##### Library process panic: "github.com/Masterminds/cookoo/safely"
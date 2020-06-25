# Logging in go

## Log
* #### Log message
    ```text
    log.Println("Logger: a message to you")
    .....Result.....
    2020/06/25 16:39:45 Logger: error panic
    ```
* #### Log error
    ```text
    log.Fatal("Logger: a error panic")
    .....Result....
    2020/06/25 16:40:38 Logger: err panic
    exit status 1
    ```
* #### Log to file
```text
logfile, _ := os.Create("./log.txt")
defer logfile.Close()
logger := log.New(logfile, "example", log.LstdFlags|log.Lshortfile)
```
##### Flags table
Flag | Description |
--- | --- |
Ldate | printing date |
Ltime | print timestamp |
Lmicrosends | add microseconds to time |
LstdFlags | both time and date |
Llongfile | printing full file path and line number |
Lshortfile | printing filename and line number |

### Logging with network resource 

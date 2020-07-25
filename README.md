# Go explained

## Noteworthy aspects of go
* ### Multiple return values
* ### A modern standard library
    * ### Modern application have common themes, such as network, crypt
    * ### Standard library: http://golang.org/pkg/
    * ### Networking and HTTP 
        * ### con, _ := net.Dial("tcp", "localhost:3306")
        * ### resp, _ := http.Get("http://google.com")
    * ### Cryptography
        * ### SHA (Secure Hash Algorithm)
        * ### TLS (Transfer layer security)
        * ### DES (Data encryption standard)
        * ### AES (Advanced Encryption standard)
        * ### HMAC (Keyed-Hash Message Authentication Code)
    * ### Data encoding, go is handled as UTF-8
* ### Concurrency with goroutines and channels 
    * ### Go is designed with parallel and concurrent processing in mind
    * ### Goroutines is managed by go runtime 
    * ### go func() { fmt.Println("hello" )}()
* ### Go the built-in toolchain 
    * ### Package management: ```go get -u ...```
    * ### Testing: ```go test```
    * ### Code Coverage:  ```go test --cover```
    * ### Formatting: ```go fmt```
    * ### Detect Race Condition: ```go run --race main.go```


## Technical
* ### [Common technical](doc/commonTechnical.md)
* ### [Web technical](doc/webTechnical.md)
* ### [Conccurency](doc/concurrency.md)
* ### [Handling error and panic](doc/errorAndPanic.md)
* ### [Logging](doc/logging.md)
* ### [Dependency Injection](doc/di.md)
* ### [Micro-services](doc/micro-services.md)

## Training project
* ### [Go basic knowledge](doc/goBasic.md)

## Common technical
#### Create server with wc
* #### TCP connection
    ```text
    wc -lk 92000
    ```
* #### UDP connection
    ```text
    wc -luk 9200
    ```
 
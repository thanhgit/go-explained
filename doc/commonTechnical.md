## Common technical

## Understanding context in go
- Using context for controlling and managing very important aspects of reliable applications
- For example: cancellation and data sharing in concurrent programming 
* ### Context with value
    - Used to share data or use request scoped values
    - When you have multiple functions and you want to share data between them -> can use contexts with context.WithValue(...) -> create context based on a parent context and adds a value to a given key
    - Key design is that everything returns a new context.Context struct in immutability, in context package 
    - Internal implementation if the context contained a map inside of it, so you can add and retrieve values by key, It can allow you store any type of data inside the context 
    ```go
    package main

    import (
        "context"
        "fmt"
    )

    func main() {
        ctx := context.Background()
        ctx = addValue(ctx)
        readValue(ctx)
    }

    func addValue(ctx context.Context) context.Context {
        return context.WithValue(ctx, "key", "test-value")
    }

    func readValue(ctx context.Context) {
        val := ctx.Value("key")
        fmt.Println(val)
    }
    ```
* ### Middlewares
    - http.Request type contains a context with scoped values throughout the HTTP pipeline 
    ```go
    package main

    import (
        "context"
        "log"
        "net/http"
        "github.com/google/uuid"
        "github.com/gorilla/mux"
    )

    func main() {
        router := mux.NewRouter()
        router.Use(guidMiddleware)
        router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)
        http.ListenAndServe(":8080", router)
    }

    func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        uuid := r.Context().Value("uuid")
        log.Printf("[%v] Returning 200 - Healthy", uuid)
        w.Write([]byte("Healthy"))
    }

    func guidMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            uuid := uuid.New()
            r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
            next.ServeHTTP(w, r)
        })
    }
    ```
* ### Context Cancellation
    - Used to propagate cancellation signal to receiver
    - How to create a cancellation context: context.WithCancel(ctx)
    ```go
    package main

    import (
        "context"
        "fmt"
        "io/ioutil"
        "net/http"
        neturl "net/url"
        "time"
    )

    func queryWithHedgedRequestsWithContext(urls []string) string {
        ch := make(chan string, len(urls))
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
        for _, url := range urls {
            go func(u string, c chan string) {
                c <- executeQueryWithContext(u, ctx)
            }(url, ch)

            select {
            case r := <-ch:
                cancel()
                return r
            case <-time.After(21 * time.Millisecond):
            }
        }

        return <-ch
    }

    func executeQueryWithContext(url string, ctx context.Context) string {
        start := time.Now()
        parsedURL, _ := neturl.Parse(url)
        req := &http.Request{URL: parsedURL}
        req = req.WithContext(ctx)

        response, err := http.DefaultClient.Do(req)

        if err != nil {
            fmt.Println(err.Error())
            return err.Error()
        }

        defer response.Body.Close()
        body, _ := ioutil.ReadAll(response.Body)
        fmt.Printf("Request time: %d ms from url%s\n", time.Since(start).Nanoseconds()/time.Millisecond.Nanoseconds(), url)
        return fmt.Sprintf("%s from %s", body, url)
    }
    ```
    - Each request is fired in a separate goroutine 
* ### ContextTimeout
    - Timeouts are a really common pattern for making external requests
    - Such as: querying a database or fetching data from another service through HTTP, gRPC
    - ctx, cancel := context.WithTimeout(ctx, time)
    - ctx, cancel := context.WithTimeout(context.Background(), 100*time,Millisecond)
    - Context in gRPC. It is used both to share data (metadata in gRPC) and to control flow (canceling a request or stream)
        - For example:

        ```go
        // metadata
        func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
            log.Printf("Sum rpc invoked with req: %v\n", req)
            md, _ := metadata.FromIncomingContext(ctx)
            log.Printf("Metadata received: %v", md)
            return &calculatorpb.SumResponse{
                Result: req.NumA + req.NumB,
            }, nil
        }
        func sum(c calculatorpb.CalculatorServiceClient) {
            req := &calculatorpb.SumRequest{
                NumA: 3,
                NumB: 10,
            }
            ctx := metadata.AppendToOutgoingContext(context.Background(), "user", "test")
            res, err := c.Sum(ctx, req)
            if err != nil {
                log.Fatalf("Error calling Sum RPC: %v", err)
            }
            log.Printf("Response: %d\n", res.Result)
        }
        ```
        
        ```go
        // Cancellation
        func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
            log.Println("Greet rpc invoked!")

            time.Sleep(500 * time.Millisecond)

            if ctx.Err() == context.Canceled {
                return nil, status.Error(codes.Canceled, "Client cancelled the request")
            }

            first := req.Greeting.FirstName
            return &greetpb.GreetResponse{
                Result: fmt.Sprintf("Hello %s", first),
            }, nil
        }
        func greetWithTimeout(c greetpb.GreetServiceClient) {
            req := &greetpb.GreetRequest{
                Greeting: &greetpb.Greeting{
                    FirstName: "Ricardo",
                    LastName:  "Linck",
                },
            }
            ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
            defer cancel()
            res, err := c.Greet(ctx, req)
            if err != nil {
                grpcErr, ok := status.FromError(err)

                if ok {
                    if grpcErr.Code() == codes.DeadlineExceeded {
                        log.Fatal("Deadline Exceeded")
                    }
                }

                log.Fatalf("Error calling Greet RPC: %v", err)
            }
            log.Printf("Response: %s\n", res.Result)
        }
        ```

    ### Read os.Stdin with channel
    ```text
    func readStdin(out chan<- []byte) {
        for {
            data := make([]byte, 1024)
            l, _ := os.Stdin.Read(data)
            if l > 0 {
                out <- data
            }
        }
    }
    ```




### Mapping file json to object
##### Library "encoding/json"
##### configuration.json
```json
{
"enabled": true,
"path": "/usr/local"
}
```
##### coding
```text
type configuration struct {
    Enabled bool    `json:"enabled"`
    Path string     `json:"path"`
}

func convert() {
    file, _ := os.Open("configuration.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    conf := configuration{}
    err := decoder.Decode(&conf)
}
```

## Mapping file yaml to object
#### library"github.com/kylelemons/go-gypsy/yaml"

#### configuration.yaml
```yaml
enabled: true
path: /usr/local
```
#### code
```text
func convert() {
    config, err := yaml.ReadFile("configuration.yaml")
    if err != nil {
    fmt.Println(err)
    }
    fmt.Println(config.Get("path"))
    fmt.Println(config.GetBool("enabled"))
}
```
```text
type configuration struct {
    Enabled bool    `yaml:"enabled"`
    Path string     `yaml:"path"`
}
func convert() {
        data, err := ioutil.ReadFile("configuration.yaml")
    	if err != nil {
    		panic(err)
    	}
    
    	errUnmarshal := yaml.Unmarshal([]byte(data), &configuration)
    
    	if errUnmarshal != nil {
    		panic(err)
    	}
}
```

### Get environment variable
```text
os.Getenv("PORT")
```


# NATS.io

## Create connection with name, timeout, health check
```text
// Set Ping Interval to 20 seconds
nc, err := nats.Connect("demo.nats.io", nats.Name("API Ping Example"), nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5))
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Do something with the connection
```
## Turn off echo in connection
```text
// Turn off echo
nc, err := nats.Connect("demo.nats.io", nats.Name("API NoEcho Example"), nats.NoEcho())
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Do something with the connection
```
## Get MaxPayload
```text
mp := nc.MaxPayload()
```

## Extra check with turn on padantic mode
```text
opts := nats.GetDefaultOptions()
opts.Url = "demo.nats.io"
// Turn on Pedantic
opts.Pedantic = true
nc, err := opts.Connect()
```

## Turn on verbose mode
```text
opts := nats.GetDefaultOptions()
opts.Url = "demo.nats.io"
// Turn on Verbose
opts.Verbose = true
nc, err := opts.Connect()
```

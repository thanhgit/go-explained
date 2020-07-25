package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	defer func() {
		if err:=recover(); err != nil {
			fmt.Printf("%s", err)
		}
	}()
	conn, err := net.Dial("tcp", "192.168.1.90:9200")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	logger := log.New(conn, "monitoring-app", log.LstdFlags)
	logger.Println("Hello, I am thanh")
	logger.Fatal("Error panic")
}

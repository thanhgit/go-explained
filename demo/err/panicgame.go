package main

import (
	"fmt"
	"net"
)

func connect() {
	var conn net.Conn
	defer func() {
		if err:=recover(); err != nil {
			fmt.Println("Error connection")
			if conn != nil {
				conn.Close()
			}
		}
	}()

	conn, err := net.Dial("tcp", "192.168.1.30:3306")
	if err != nil {
		panic("Error connect")
	}

	//con.Close()
}
func main()  {
	connect()
}

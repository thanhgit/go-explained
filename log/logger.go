package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()
	logger := log.New(logfile, "logger-app ", log.LstdFlags)
	logger.Println("Logger: error panic") // 2020/06/25 16:39:45 Logger: error panic
	logger.Fatal("Logger: err panic") // 2020/06/25 16:40:38 Logger: err panic \n exit status 1
}

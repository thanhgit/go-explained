package main

import (
	"log"
	"log/syslog"
	"os"
)

func syslogVersion1() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	syslogger, err := syslog.NewLogger(priority, flags)
	syslogger.SetOutput(os.Stdout)
	if err != nil {
		panic(err)
	}
	syslogger.Println("Hello")
}
func main() {
	syslogger, err := syslog.New(syslog.LOG_INFO, "syslog_example")
	defer log.SetOutput(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(syslogger)
	log.Println("Log entry")

}

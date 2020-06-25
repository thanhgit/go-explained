package main

import (
	"log"
	"log/syslog"
)

func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	syslogger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		panic(err)
	}

	syslogger.Println("This is syslog")

}

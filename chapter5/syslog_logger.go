package main

import (
	"log/syslog"
)

func main() {
	// priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	// flags := log.Ldate | log.Lshortfile
	// logger, err := syslog.NewLogger(priority, flags)
	logger, err := syslog.New(syslog.LOG_LOCAL3, "junfeng")
	if err != nil {
		// fmt.Printf("Can't attach to syslog: %s", err)
		// return
		panic("Cannot attach to syslog")
	}
	defer logger.Close()

	logger.Debug("Debug message.")
	logger.Notice("Notice message.")
	logger.Warning("Warning message.")
	logger.Alert("Alert message.")
}

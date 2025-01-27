package main

import "fmt"

/*
Imagine you're building a logging system for an application. You want to
ensure that only one instance of the logger is used throughout the entire
application, so you avoid opening multiple log files or creating multiple
connections to the logging service.
*/

var logger *Logger

type Logger struct {
}

func NewLogger() *Logger {
	if logger == nil {
		logger = &Logger{}
	}

	return logger
}

func main() {
	logger1 := NewLogger()
	fmt.Printf("Logger is %v\n", &logger1)

	logger2 := NewLogger()
	fmt.Printf("Logger2 is %v\n", &logger2)

	if logger1 == logger2 {
		fmt.Println("Both loggers are the same instance!")
	} else {
		fmt.Println("Loggers are different instances.")
	}
}

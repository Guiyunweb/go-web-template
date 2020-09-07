package log

import (
	"fmt"
	"time"
)

type Logger struct {
	level string
	meg   string
}

const (
	PANIC   = "| PANIC   |"
	ERROR   = "| ERROR   |"
	WARNING = "| WARNING |"
	INFO    = "| INFO    |"
	DEBUG   = "| DEBUG   |"
)

func Println(logger Logger) {
	fmt.Printf("[system] %s %s %s \n", time.Now().Format("2006-01-02 15:04:05"), logger.level, logger.meg)
}

func Panic(msg string) {
	logger := Logger{
		PANIC,
		msg,
	}
	Println(logger)
}

func Error(msg string) {
	logger := Logger{
		ERROR,
		msg,
	}
	Println(logger)
}

func Waring(msg string) {
	logger := Logger{
		WARNING,
		msg,
	}
	Println(logger)
}

func Info(msg string) {
	logger := Logger{
		INFO,
		msg,
	}
	Println(logger)
}

func Debug(msg string) {
	logger := Logger{
		DEBUG,
		msg,
	}
	Println(logger)
}

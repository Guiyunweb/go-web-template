package log

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	level string
	meg   string
}

const (
	_PANIC   = "| PANIC   |"
	_ERROR   = "| ERROR   |"
	_WARNING = "| WARNING |"
	_INFO    = "| INFO    |"
	_DEBUG   = "| DEBUG   |"
)

func _Println(logger Logger) {
	fmt.Printf("[system] %s %s %s \n", time.Now().Format("2006-01-02 15:04:05"), logger.level, logger.meg)
}

func Panic(msg string) {
	logger := Logger{
		_PANIC,
		msg,
	}
	_Println(logger)
	os.Exit(0)
}

func Error(msg string) {
	logger := Logger{
		_ERROR,
		msg,
	}
	_Println(logger)
}

func Waring(msg string) {
	logger := Logger{
		_WARNING,
		msg,
	}
	_Println(logger)
}

func Info(msg string) {
	logger := Logger{
		_INFO,
		msg,
	}
	_Println(logger)
}

func Debug(msg string) {
	logger := Logger{
		_DEBUG,
		msg,
	}
	_Println(logger)
}

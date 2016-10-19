package api_utils

import (
	"log"
	"os"
)

type Loggable interface {
	CollectPanic(lw LogWriterStrategy)
}

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger)CollectPanic(lw LogWriterStrategy) {
	if r := recover(); r != nil {
		lw(r)
	}
}

type LogWriterStrategy func(v interface{})

func FileSystemWriter(v interface{}) {
	f, err := os.OpenFile("resources/testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(v)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(v)
}

func ConsoleWriter(v interface{}) {
	log.Println(v)
}




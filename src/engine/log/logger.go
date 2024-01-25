package log

import "fmt"

type Logger struct {
}

var instance *Logger

func GetLogger() *Logger {
	if instance == nil {
		instance = &Logger{}
	}
	return instance
}

func (l *Logger) Message(m string) {
	fmt.Println(m)
}

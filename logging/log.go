package logging

import (
	"fmt"
	"time"
)

const (
	LOG_INFO   = iota
	LOG_ERROR  = iota
	LOG_RESULT = iota
)

type Log struct {
	entries []*Entry
}

type Entry struct {
	kind    int
	time    string
	message string
}

func (l *Log) Info(content interface{}) {
	l.entries = append(l.entries, &Entry{kind: LOG_INFO, time: time.Now().Format(time.RFC3339), message: fmt.Sprintf("%v", content)})
}

func (l *Log) Result(content interface{}) {
	l.entries = append(l.entries, &Entry{kind: LOG_RESULT, time: time.Now().Format(time.RFC3339), message: fmt.Sprintf("%v", content)})
}

func (l *Log) Error(content interface{}) {
	l.entries = append(l.entries, &Entry{kind: LOG_ERROR, time: time.Now().Format(time.RFC3339), message: fmt.Sprintf("%v", content)})
}

func (l *Log) Flush() {
	for _, entry := range l.entries {
		if entry.kind == LOG_ERROR {
			fmt.Println("!! " + entry.message)
		}
	}

	fmt.Println()
	fmt.Println("Info: ")

	for _, entry := range l.entries {
		if entry.kind == LOG_INFO {
			fmt.Println("* " + entry.message)
		}
	}

	fmt.Println()
	fmt.Println("Result: ")

	for _, entry := range l.entries {
		if entry.kind == LOG_RESULT {
			fmt.Println("* " + entry.message)
		}
	}
}

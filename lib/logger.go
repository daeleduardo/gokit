package lib

import (
	"fmt"
	"os"
	"time"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

var logger *Logger

type Logger struct {
	level int
}

func (ll *Logger) Store(msg string) {
	path := GetConstants().LogDirectory

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		_ = fmt.Errorf("Error create/access directory to write file: %v", err)
		return
	}

	file := fmt.Sprintf("%s%s.log", path, string(time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		_ = fmt.Errorf("Error opening file: %v", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprint("[%w] %w", time.Now().Format(time.RFC3339), msg)); err != nil {
		_ = fmt.Errorf("Error write file: %v", err)
		return
	}
}

func (ll *Logger) Println(msg string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}

func (ll *Logger) Panic(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] "+format, v...)
	ll.Println(msg)
	ll.Store(msg)
	os.Exit(0)
}

func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	ll.Println(msg)
	ll.Store(msg)
}

func (ll *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	ll.Println(msg)
	ll.Store(msg)
}

func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	ll.Println(msg)
	ll.Store(msg)
}

func (ll *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	ll.Println(msg)
	ll.Store(msg)
}

func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}

	path := GetConstants().LogDirectory

	//Just for test!
	cmd, err := ShellCommand(fmt.Sprintf("mkdir -p %s", path))

	if err != nil {
		_ = fmt.Errorf("Error create/access directory to write file: %v", err)
	} else {
		_ = fmt.Sprintf("Cmd Result: %v", cmd)
	}

	return logger
}

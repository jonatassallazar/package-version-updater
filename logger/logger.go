package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	Level uint8
}

// ANSI color codes - https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	FGBlack  string = "30"
	FGRed    string = "31"
	FGGreen  string = "32"
	FGYellow string = "33"
	FGBlue   string = "34"
	FGPurple string = "35"
	FGCyan   string = "36"
	FGWhite  string = "37"
	BGBlack  string = "40"
	BGRed    string = "41"
	BGGreen  string = "42"
	BGYellow string = "43"
	BGBlue   string = "44"
	BGPurple string = "45"
	BGCyan   string = "46"
	BGWhite  string = "47"
)
const (
	DEBUG uint8 = 5
	INFO  uint8 = 4
	WARN  uint8 = 3
	ERROR uint8 = 2
	FATAL uint8 = 1
	PANIC uint8 = 0
)

func (l *Logger) message(level, color, message string) string {
	return fmt.Sprintf("\033[%sm %s | [%s] %s \033[0m", color, level, time.Now().Format(time.RFC1123Z), message)
}

func (l *Logger) rawMessage(level, message string) string {
	return fmt.Sprintf("\n%s | [%s] %s", level, time.Now().Format(time.RFC1123Z), message)
}

func (l *Logger) BreakLine() {
	l.writeToLogFile("\n")
	fmt.Println()
}

func (l *Logger) Debug(message string) {
	if l.Level >= DEBUG {

		color := "0;" + FGGreen
		level := "DEBUG"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))
	}
}

func (l *Logger) Log(message string) {
	if l.Level >= INFO {

		color := "0;" + FGWhite
		level := "INFO"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))
	}
}

func (l *Logger) Warning(message string) {
	if l.Level >= WARN {

		color := "0;" + FGYellow
		level := "WARNING"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))
	}
}

func (l *Logger) Error(message string) {
	if l.Level >= ERROR {

		color := "0;" + FGRed
		level := "ERROR"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))
	}
}

func (l *Logger) Fatal(message string) {
	if l.Level >= FATAL {

		color := FGBlack + ";" + BGRed
		level := "FATAL"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))
	}
}

func (l *Logger) Panic(message string) {
	if l.Level >= PANIC {

		color := FGRed + ";" + BGWhite
		level := "PANIC"

		l.writeToLogFile(l.rawMessage(level, message))

		fmt.Println(l.message(level, color, message))

		panic(message)
	}
}

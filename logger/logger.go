package logger

import (
	"fmt"
	"time"
)

type Logger struct {
}

// ANSI color codes - https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	FGBlack  = "30"
	FGRed    = "31"
	FGGreen  = "32"
	FGYellow = "33"
	FGBlue   = "34"
	FGPurple = "35"
	FGCyan   = "36"
	FGWhite  = "37"
	BGBlack  = "40"
	BGRed    = "41"
	BGGreen  = "42"
	BGYellow = "43"
	BGBlue   = "44"
	BGPurple = "45"
	BGCyan   = "46"
	BGWhite  = "47"
)

func (l *Logger) message(color, message string) string {
	return fmt.Sprintf("\033[%sm[%s] %s \033[0m", color, time.Now().Format(time.RFC1123Z), message)
}

func (l *Logger) rawMessage(message string) string {
	return fmt.Sprintf("\n[%s] %s", time.Now().Format(time.RFC1123Z), message)
}

func (l *Logger) BreakLine() {
	l.writeToLogFile("\n")
	fmt.Println()
}

func (l *Logger) Verbose(message string) {
	color := "0;" + FGGreen

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

func (l *Logger) Log(message string) {
	color := "0;" + FGWhite

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

func (l *Logger) Info(message string) {
	color := "0;" + FGBlue

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

func (l *Logger) Error(message string) {
	color := "0;" + FGRed

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

func (l *Logger) Fatal(message string) {
	color := FGBlack + ";" + BGRed

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

func (l *Logger) Panic(message string) {
	color := FGRed + ";" + BGWhite

	l.writeToLogFile(l.rawMessage(message))

	fmt.Println(l.message(color, message))
}

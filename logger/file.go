package logger

import "os"

func (l *Logger) writeToLogFile(message string) {
	file, err := os.OpenFile("./debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		l.Error(err.Error())
	}
	defer file.Close()

	file.Write([]byte(message))
}

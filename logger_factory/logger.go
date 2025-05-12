package logger_factory

import (
	"log"
	"os"
)

const filePermission = 0666
const jlangLs = "[JLANG LS]"

func GetLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, filePermission)

	if err != nil {
		panic("No can do for logging")
	}

	return log.New(logfile, jlangLs, log.Ldate|log.Ltime|log.Lshortfile)
}

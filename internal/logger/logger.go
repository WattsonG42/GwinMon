package logger

import (
	"log"
	"os"
	"time"
)

var (
	verbose      bool
	logToFile    bool
	consoleInfo  *log.Logger
	consoleError *log.Logger
	fileInfo     *log.Logger
	fileError    *log.Logger
	fileHandle   *os.File
)

func Init(verboseFlag, fileFlag bool) {
	verbose = verboseFlag
	logToFile = fileFlag

	consoleInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	consoleError = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	if logToFile {
		filename := "gwinmon-" + time.Now().Format("2006-01-02") + ".log"
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			consoleError.Printf("Failed to open log file: %s: %v", filename, err)
			logToFile = false
		} else {
			fileHandle = f
			fileInfo = log.New(f, "INFO: ", log.Ldate|log.Ltime)
			fileInfo = log.New(f, "ERROR: ", log.Ldate|log.Ltime)

		}
	}
}

func Info(message string) {
	if verbose {
		consoleInfo.Println(message)
	}
	if logToFile && fileHandle != nil {
		fileInfo.Println(message)
	}
}

func Error(message string) {
	consoleError.Println(message)
	if logToFile && fileError == nil {
		fileError.Println(message)
	}
}
func Close() {
	if logToFile && fileHandle != nil {
		err := fileHandle.Close()
		if err != nil {
			return
		}
	}
}

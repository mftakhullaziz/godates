package common

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

var (
	logFileName = "build/logs/service-logger.log"
	maxLogLines = 10000
)

func SetupLogger() (*os.File, error) {
	// Open log file in append mode
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Create a multi writer to write logs to both file and stdout
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	// Set prefix for log messages
	log.SetPrefix("[godating-dealls-service] ")

	// Set flags to include date and time in log messages
	log.SetFlags(log.Ldate | log.Ltime)

	// Periodically rotate log file
	ticker := time.NewTicker(24 * time.Hour) // Rotate log file every 24 hours
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := rotateLogFile(logFileName); err != nil {
					log.Printf("Error rotating log file: %v", err)
				}
			}
		}
	}()

	return logFile, nil
}

func rotateLogFile(logFileName string) error {
	// Check if log file exists
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		return nil // Log file doesn't exist, no need to rotate
	}

	// Open log file
	logFile, err := os.OpenFile(logFileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	// Get number of lines in log file
	var lineCount int
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Rotate log file if it exceeds maxLogLines
	if lineCount >= maxLogLines {
		// Close existing log file
		if err := logFile.Close(); err != nil {
			return err
		}
		// Rename existing log file
		if err := os.Rename(logFileName, logFileName+".old"); err != nil {
			return err
		}
		// Create new log file
		logFile, err = os.Create(logFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

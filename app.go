package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	logger := log.New(log.Writer(), "", log.LstdFlags)

	for i := 0; i < 20; i++ {
		// Generate a random log message
		message := randomLogMessage()

		// Format the log message
		formattedMessage := formatLogMessage(message)

		// Write the formatted log message to stdout
		logger.Println(formattedMessage)

		// Wait for a random amount of time before generating the next log message
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func randomLogMessage() string {
	// List of possible log messages
	messages := []string{
		"Something happened",
		"An error occurred",
		"User logged in",
		"File saved successfully",
		"Database connection established",
		"Message sent",
		"Task completed",
		"Server started",
		"Configuration updated",
		"API endpoint hit",
	}

	// Pick a random message from the list
	return messages[rand.Intn(len(messages))]
}

func formatLogMessage(message string) string {
	// Get the current time
	currentTime := time.Now()

	// Format the log message
	return fmt.Sprintf("[%s] %s", currentTime.Format("2006-01-02 15:04:05"), message)
}

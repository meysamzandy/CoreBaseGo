package main

import (
	"fmt"
	"time"
)

func main() {
	// Initialize the worker (this could involve setting up dependencies, etc.)
	fmt.Println("Starting hourly worker...")

	// Create a ticker that ticks every hour
	ticker := time.NewTicker(1 * time.Second) // Ticks every hour
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Perform the task that you want to run every hour
			processHourlyTask()
		}
	}
}

// processHourlyTask simulates a task that runs every hour
func processHourlyTask() {
	// Example task: print a log message
	fmt.Println("Running hourly task...")

	// Simulate task processing (replace with your actual logic)
	time.Sleep(5 * time.Second) // Simulate some processing
	fmt.Println("Hourly task completed!")
}

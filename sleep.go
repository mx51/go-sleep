package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	sleepEnvVarName = "PRESTOP_SLEEP"
)

var (
	// Default - no sleep
	defaultSleepSeconds = 0
)

func main() {
	// Log useful start message
	log.Println("PRESTOP SLEEP: STARTING")

	// Determine sleep time duration
	sleepSeconds := checkSleepEnvVar()

	// Sleep for duration
	if sleepSeconds > 0 {
		log.Printf("Sleeping for %d seconds ...", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

	// Log useful end message
	log.Println("PRESTOP SLEEP: COMPLETE")
}

func checkSleepEnvVar() int {
	sleepSeconds := defaultSleepSeconds
	// Read env var or default
	if sleepEnvVarValue := os.Getenv(sleepEnvVarName); sleepEnvVarValue != "" {
		// Parse env var value
		i, err := strconv.Atoi(sleepEnvVarValue)
		if err != nil {
			log.Printf("error: could not parse env var %s: %v", sleepEnvVarName, err)
		} else {
			sleepSeconds = i
		}
	}
	return sleepSeconds
}

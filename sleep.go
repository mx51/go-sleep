package main

import (
	"flag"
	"log"
	"os"
	"time"
)

const (
	sleepEnvVarName = "SLEEP_SECONDS"
)

var (
	sleepSeconds int
)

func init() {
	flag.IntVar(&sleepSeconds, "s", 0, "sleep duration in seconds")
}

func main() {
	// Process cli args
	flag.Parse()

	// Lookup HOSTNAME (TODO - use better logging)
	hostname := lookupHostname()

	// Log useful start message
	log.Printf("instance=%s - SLEEP STARTING", hostname)

	// Sleep for duration
	if sleepSeconds > 0 {
		log.Printf("instance=%s - sleeping for %d seconds ...", hostname, sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}

	// Log useful end message
	log.Printf("instance=%s - SLEEP COMPLETE", hostname)
}

func lookupHostname() string {
	if hostname := os.Getenv("HOSTNAME"); hostname != "" {
		return hostname
	}
	return "unset"
}

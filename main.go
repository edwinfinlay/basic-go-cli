package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	url      string
	logFile  string
	logger   *log.Logger
	interval time.Duration
)

func main() {
	flag.StringVar(&url, "url", "", "URL to check")
	flag.StringVar(&logFile, "logfile", "healthcheck.log", "File to log output to")
	flag.DurationVar(&interval, "interval", 2*time.Second, "Interval between health checks")
	flag.Parse()

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open or create log file %s: %v\n", logFile, err)
		log.Fatal(err)
	}
	defer file.Close()
	logger = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	if url != "" {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			checkURL(url)
		}
	}
}

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting url %s: %v", url, err)
		return
	}
	logger.Printf("Checked url: %s, Status: %d", url, resp.StatusCode)
}

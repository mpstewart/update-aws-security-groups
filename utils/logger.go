package utils

import (
	"log"
	"os"
)

// Logger - singleton access to a logger which will print long format output
var Logger *log.Logger = log.New(
	os.Stdout,
	"",
	log.Lshortfile,
)

package logger

import (
	"altComparator/pkg/IO"
	"log"
	"os"
)

func NewLogger(typeLogger string) *log.Logger {
	f := os.Stderr
	if IO.Debug {
		f = os.Stdout
	}
	if typeLogger == "ERROR" {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return log.New(f, typeLogger+": ", log.Ldate|log.Ltime)
	}
}

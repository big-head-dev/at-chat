package main

import (
	"time"
)

func getTimeNow(format string) string {
	now := time.Now()
	return now.Format(format)
}

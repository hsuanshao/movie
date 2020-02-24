package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	currentUsers       int   = 0
	processedCounter   int   = 0
	currentRequestRate int   = 0
	firstTimeMs        int64 = timeMs(time.Now())
	lastTimeMs         int64 = 0
)

// MovieInfo describe request result and current request user number & request rate usage status
type MovieInfo struct {
	Message            string
	MovieJSON          *[]byte
	CurrentUsers       int
	CurrentRequestRate int
}

func timeMs(t time.Time) int64 {
	return t.UnixNano() / time.Millisecond.Nanoseconds()
}

func main() {
	fmt.Printf("Start to listen user request.....")
	go func() {
		// HTTP monitor
		http.HandleFunc("/", monitor)
		http.ListenAndServe(":8083", nil)
	}()

	Listen()

}

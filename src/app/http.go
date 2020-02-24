package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// HTTPListener response to listen port 8083 and provide monitor result
func HTTPListener() {
	http.HandleFunc("/", monitor)
	http.ListenAndServe(":8083", nil)
}

func monitor(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Current connection: %s \n", strconv.Itoa(currentUsers))
	fmt.Fprintf(w, "Current Request rate: %s \n", strconv.Itoa(currentRequestRate))
	fmt.Fprintf(w, "Processed request: %s \n", strconv.Itoa(processedCounter))
	pJobs := processedCounter
	if processedCounter > 30 {
		pJobs = processedCounter % 30
	}
	fmt.Fprintf(w, "Remaing jobs: %s", strconv.Itoa(currentRequestRate-pJobs))
}

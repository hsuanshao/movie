package main

/**
  Date: 2020-02-21
  Author: William Chang
  Description: this is a TCP application to query movie information
  Port: 8081
  Monitor: 127.0.0.1:8083
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

// omdb authorization free test key, it has limitation in 1000 query/per day
const (
	authorization string = "cb132a69"
	rateLimit     int    = 30
)

var (
	currentUsers       int   = 0
	processedCounter   int   = 0
	currentRequestRate int   = 0
	firstTimeMs        int64 = timeMs(time.Now())
	lastTimeMs         int64 = 0
)

func main() {
	// TCP
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("listener error:", err)
	}

	go func() {
		// HTTP monitor
		http.HandleFunc("/", monitor)
		http.ListenAndServe(":8083", nil)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
		}
		currentUsers++
		go connHandler(conn)
	}
}

func timeMs(t time.Time) int64 {
	return t.UnixNano() / time.Millisecond.Nanoseconds()
}

func connHandler(conn net.Conn) {
	// if user did not request anything, it will disconnect in 20 seconds
	timeoutDuration := 20 * time.Second
	conn.SetDeadline(time.Now().Add(timeoutDuration))
	conn.Write([]byte("Search Movie:"))
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		currentUsers--
		log.Println("user left..")
		conn.Close()
		return
	}

	// Get client query string
	moviename := string(bufferBytes)
	clientAddr := conn.RemoteAddr().String()
	response := fmt.Sprintf(moviename + " from " + clientAddr + "\n")

	// Get movie information by user input
	// Dataset is omdb opendata
	result, err := askMovie(moviename)
	if err != nil {
		log.Printf("set api query failure")
	}
	conn.Write([]byte("Movie Name: " + response))
	conn.Write(result)
	conn.Write([]byte(fmt.Sprintf("\n\n")))
	currentRequestRate++

	// rate limit
	lastTimeMs = timeMs(time.Now())
	duration := lastTimeMs - firstTimeMs

	if currentRequestRate == rateLimit && duration <= 1000 {
		firstTimeMs = timeMs(time.Now())
		log.Println("Hit rate limit")
		conn.Write([]byte("Hit rate limit"))
		return
	}

	connHandler(conn)
}

func askMovie(moviename string) ([]byte, error) {
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		log.Printf("request " + moviename + " info failure")
		return nil, fmt.Errorf("can't set request")
	}

	q := req.URL.Query()
	// OMDB open api request parameter
	q.Add("apikey", authorization)
	q.Add("t", moviename)

	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print("get movie info failure", err)
		return nil, fmt.Errorf("API request failure")
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("read response failure")
		return nil, fmt.Errorf("Get API response failure")
	}
	processedCounter++
	return responseData, nil
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

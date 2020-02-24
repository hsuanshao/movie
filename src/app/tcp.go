package main

import (
	"bufio"
	"fmt"
	"log"
	"movie"
	"net"
	"time"
)

const rateLimit int = 30

// Listen response listen port 8081 request
func Listen() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("listener error:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
		}
		currentUsers++
		go connHandler(conn)
	}
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
	result, err := movie.AskMovie(moviename)
	if err != nil {
		log.Printf("set api query failure")
	}
	processedCounter++
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

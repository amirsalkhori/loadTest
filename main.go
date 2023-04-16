package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	// Set up the load test parameters
	numRequests := 100000
	// concurrencyLevel := 100
	timeout := time.Duration(5 * time.Second)

	// Create a wait group to ensure all goroutines have finished
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Create a client with a timeout
	client := http.Client{
		Timeout: timeout,
	}

	// Start the load test
	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		// Start a goroutine for each request
		go func() {
			defer wg.Done()

			url := "https://your_domain"
			method := "POST"

			payload := strings.NewReader("code=your payload")

			req, err := http.NewRequest(method, url, payload)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("cache-control", "no-cache")
			req.Header.Add("content-type", "application/x-www-form-urlencoded")
			req.Header.Add("postman-token", "token_postman")
			req.Header.Add("acce", "application/json")
			req.Header.Add("host", "your_host")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("")
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Calculate the load test duration and requests per second
	duration := time.Since(startTime)
	requestsPerSecond := float64(numRequests) / duration.Seconds()

	fmt.Printf("Load test completed in %v with %d requests per second\n", duration, int(requestsPerSecond))
}

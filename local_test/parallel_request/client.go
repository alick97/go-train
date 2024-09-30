package parallel_request

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func makeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close() // Important: Close the response body

	data, _ := io.ReadAll(resp.Body)
	elapsed := time.Since(start)
	ch <- fmt.Sprintf("Fetched %s in %s, data: %s", url, elapsed, string(data))
}

func RunParallelExample() {
	urlTemplate := "http://127.0.0.1:9999?v=%d"
	urls := make([]string, 10000)
	for i := range urls {
		urls[i] = fmt.Sprintf(urlTemplate, i)
	}

	ch := make(chan string, 100) // Create a channel to receive results
	defer close(ch)              // Close the channel when done

	for _, url := range urls {
		go makeRequest(url, ch) // Launch a goroutine for each request
	}

	for range urls { // Receive results from all goroutines
		fmt.Println(<-ch)
	}
}

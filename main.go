package main

import (
	"fmt"
	"os"
	"sync"
	"test-http-status-checker/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url_file>")
		os.Exit(1)
	}

	urlFile := os.Args[1]
	urls, err := utils.ReadURLs(urlFile)
	if err != nil {
		fmt.Printf("Error reading URL file: %s\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go utils.CheckURL(url, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}

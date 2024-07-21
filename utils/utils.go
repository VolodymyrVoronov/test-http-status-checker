package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	Reset = "\033[0m"
	Green = "\033[32m"
	Red   = "\033[31m"
)

func CheckURL(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("%sERROR: %s => %s%s", Red, url, err, Reset)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		results <- fmt.Sprintf("%sURL: %s => Status: %d%s", Green, url, resp.StatusCode, Reset)
	} else {
		results <- fmt.Sprintf("%sURL: %s => Status: %d%s", Red, url, resp.StatusCode, Reset)
	}
}

func ReadURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

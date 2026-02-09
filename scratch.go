package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func countWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func main() {
	resp, _ := http.Get("https://mistwire.com")
	result := countWords(resp.Body)
	fmt.Println(result)

	s := strings.NewReader("This is just five words... wait")
	result = countWords(s)
	fmt.Println(result)
}

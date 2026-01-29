package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// open a file
	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error %s: ", err)
	}
	defer f.Close()

	// the frequency of the words in the file
	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from the freq in main: %s", err)
	}

	// Display the word frequencies

	// Sort the word frequencies
	pairs := sortWordFrequency(words)

	// print the sorted results
	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n, pair.Key, pair.Value")
	}

	// word with greatest freq & it's freq
	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord: %s\n", err)
	}
	fmt.Printf("%#v has a frequency of %d", w, n)
}

func freq(r io.Reader) (map[string]int, error) {
	// create a map to store word frequencies
	wordFreq := make(map[string]int)

	// create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	// read each word and update the word frequencies
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return wordFreq, nil
}

// for sorting frequency of words
type Pair struct {
	Key   string
	Value int
}

func sortWordFrequency() {}

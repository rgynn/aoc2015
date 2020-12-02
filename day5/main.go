package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	one := fineNiceOne(file)
	two := fineNiceTwo(file)

	log.Printf("Answer to exercise one: %d", one)
	log.Printf("Answer to exercise two: %d", two)
}

func fineNiceOne(file []byte) int {

	var entries []entryOne

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		entries = append(entries, entryOne{Line: scanner.Text()})
	}

	var count int
	for _, e := range entries {
		if e.Nice() {
			count++
		}
	}

	return count
}

func fineNiceTwo(file []byte) int {

	var entries []entryTwo

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		entries = append(entries, entryTwo{Line: scanner.Text()})
	}

	var count int
	for _, e := range entries {
		if e.Nice() {
			count++
		}
	}

	return count
}

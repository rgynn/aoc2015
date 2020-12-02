package main

import (
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	input := string(file)

	var one int
	for _, instruction := range input {
		switch instruction {
		case '(':
			one++
		case ')':
			one--
		}
	}

	var two int
	for i, instruction := range input {
		switch instruction {
		case '(':
			two++
		case ')':
			two--
		}
		if two == -1 {
			two = i + 1
			break
		}
	}

	log.Printf("Answer to exercise one: \n")
	log.Printf("Santa ends up on floor: %d", one)

	log.Printf("Answer to exercise two: \n")
	log.Printf("Santa ends up on floor: %d", two)
}

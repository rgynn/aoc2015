package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

/*
 * 123 -> x means that the signal 123 is provided to wire x.
 * x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
 * p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
 * NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.
 */

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	circuit := map[string][2]byte{}

	gateRegexp := regexp.MustCompile("[A-Z]+")
	outputRegexp := regexp.MustCompile("-> (.+)")
	inputRegexp := regexp.MustCompile("([a-z0-9]+) ")
	wireRegexp := regexp.MustCompile("[a-z]+")

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {

		line := scanner.Text()

		signal := [2]byte{}
		if strings.Contains(line, "AND") {
			linesplit := strings.Split(line, " AND ")
			linesplit2 := strings.Split(linesplit[1], " -> ")[0]
			from1 := linesplit[0]
			from2 := linesplit2
		}

		circuit[line[:2]] = signal
	}
}

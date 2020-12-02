package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type box struct {
	l, w, h int
}

func (b box) neededPaper() (sum int) {

	lSide := 2 * b.l * b.w
	wSide := 2 * b.w * b.h
	topBottom := 2 * b.h * b.l

	sum = lSide + wSide + topBottom

	// Get the smallest side, and add to sum
	dimensions := []int{lSide / 2, wSide / 2, topBottom / 2}

	sort.Slice(dimensions, func(i, j int) bool {
		return dimensions[i] < dimensions[j]
	})

	sum += dimensions[0]

	return sum
}

func (b box) neededRibbon() int {

	dimensions := []int{b.l, b.w, b.h}

	sort.Slice(dimensions, func(i, j int) bool {
		return dimensions[i] < dimensions[j]
	})

	bow := dimensions[0] * dimensions[1] * dimensions[2]

	return bow + dimensions[0] + dimensions[0] + dimensions[1] + dimensions[1]
}

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	boxes := []box{}
	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)

		split := strings.Split(line, "x")

		l, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("FATAL: %v", err)
		}

		w, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("FATAL: %v", err)
		}

		h, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatalf("FATAL: %v", err)
		}

		boxes = append(boxes, box{l, w, h})
	}

	var paper, ribbon int
	for _, box := range boxes {
		paper += box.neededPaper()
		ribbon += box.neededRibbon()
	}

	log.Printf("Needed paper: %d \n", paper)
	log.Printf("Needed ribbon: %d \n", ribbon)
}

package main

import (
	"io/ioutil"
	"log"
)

type position struct {
	name string
	x, y int
}

type house struct {
	x, y, visits int
}

func main() {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	one := firstRun(file)
	two := secondRun(file)

	log.Printf("First run santa goes through %d houses", len(one))
	log.Printf("Second run santa and robo santa goes through %d houses", len(two))
}

func firstRun(file []byte) []house {

	// santa coordinates
	santa := position{name: "santa"}

	// list of visited houses
	houses := []house{
		house{visits: 1}, // starting house
	}

	for _, instruction := range string(file) {

		switch instruction {
		case '^':
			santa.y++
		case 'v':
			santa.y--
		case '<':
			santa.x++
		case '>':
			santa.x--
		}

		var visited bool
		for i, h := range houses {
			if h.x == santa.x && h.y == santa.y {
				visited = true
				houses[i].visits++
				break
			}
		}

		if !visited {
			houses = append(houses, house{x: santa.x, y: santa.y, visits: 1})
		}
	}

	return houses
}

func secondRun(file []byte) []house {

	santas := []position{
		position{
			name: "santa",
		},
		position{
			name: "robo santa",
		},
	}

	// list of visited houses
	houses := []house{
		house{visits: 2},
	}

	turn := len(santas)

	for _, instruction := range string(file) {

		// switch turns
		turn = turn % len(santas)

		switch instruction {
		case '^':
			santas[turn].y++
		case 'v':
			santas[turn].y--
		case '<':
			santas[turn].x++
		case '>':
			santas[turn].x--
		}

		var visited bool
		for i, h := range houses {
			if h.x == santas[turn].x && h.y == santas[turn].y {
				visited = true
				houses[i].visits++ // house already visited, increment visits
				break
			}
		}

		// house not visited yet, add to slice
		if !visited {
			houses = append(houses, house{x: santas[turn].x, y: santas[turn].y, visits: 1})
		}

		// make sure to increment to switch turns
		turn++
	}

	return houses
}

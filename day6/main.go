package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type position struct {
	x, y int
}

type instruction struct {
	method   string
	from, to position
}

func (i *instruction) Go(lights [1000][1000]light) [1000][1000]light {
	for y, row := range lights {
		for x := range row {
			if y >= i.from.y && y <= i.to.y &&
				x >= i.from.x && x <= i.to.x {
				switch i.method {
				case "turn on":
					lights[y][x].TurnOn()
				case "turn off":
					lights[y][x].TurnOff()
				case "toggle":
					lights[y][x].Toggle()
				}
			}
		}
	}
	return lights
}

type light struct {
	on         bool
	brightness int
}

func (l *light) TurnOn() {
	l.on = true
	l.brightness++
}

func (l *light) TurnOff() {
	l.on = false
	if l.brightness-1 <= 0 {
		l.brightness = 0
	} else {
		l.brightness--
	}
}

func (l *light) Toggle() {
	l.on = !l.on
	l.brightness += 2
}

func main() {

	instructions, err := readInput()
	if err != nil {
		log.Fatalf("FATAL: %v", err)
	}

	lights := [1000][1000]light{}

	for _, i := range instructions {
		lights = i.Go(lights)
	}

	var one int
	for _, row := range lights {
		for _, l := range row {
			if l.on {
				one++
			}
		}
	}

	var two int
	for _, row := range lights {
		for _, l := range row {
			two += l.brightness
		}
	}

	log.Printf("Answer to exercise one: %d", one)
	log.Printf("Answer to exercise two: %d", two)
}

func readInput() ([]instruction, error) {

	file, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}

	var instructions []instruction

	instructionRegex := regexp.MustCompile(`\b(turn off|turn on|toggle) \b(\d{1,3}),(\d{1,3}) through \b(\d{1,3}),(\d{1,3})`)

	scanner := bufio.NewScanner(bytes.NewBufferString(string(file)))
	for scanner.Scan() {

		line := scanner.Text()
		match := instructionRegex.FindAllStringSubmatch(line, 5)

		method := match[0][1]

		fromX, err := strconv.Atoi(match[0][2])
		if err != nil {
			return nil, err
		}

		fromY, err := strconv.Atoi(match[0][3])
		if err != nil {
			return nil, err
		}

		toX, err := strconv.Atoi(match[0][4])
		if err != nil {
			return nil, err
		}

		toY, err := strconv.Atoi(match[0][5])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction{
			method: method,
			from:   position{x: fromX, y: fromY},
			to:     position{x: toX, y: toY},
		})
	}

	return instructions, nil
}

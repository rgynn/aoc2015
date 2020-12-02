package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"strings"
)

const input = "bgvyzdsv"

func main() {

	one := findHash("00000")
	two := findHash("000000")

	log.Println("--------")
	log.Printf("Answer to part one: %d \n", one)
	log.Println("--------")
	log.Printf("Answer to part two: %d \n", two)
}

func findHash(prefix string) int {

	var decimal int
	for {

		key := fmt.Sprintf("%s%d", input, decimal)

		md5hasher := md5.New()

		_, err := io.WriteString(md5hasher, key)
		if err != nil {
			log.Fatalf("FATAL: %v", err)
		}

		hex := fmt.Sprintf("%x", md5hasher.Sum(nil))

		if strings.HasPrefix(hex, prefix) {

			log.Println("--------")
			log.Printf("Decimal: %d\n", decimal)
			log.Printf("Key: %s\n", key)
			log.Printf("Hex: %s\n", hex)

			return decimal
		}

		decimal++
	}
}

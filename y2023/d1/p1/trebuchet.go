package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		log.Panicln("invalid number of arguments passed, pass in the calibation document file path")
	}

	path := os.Args[1]
	lines, err := readDocument(path)
	if err != nil {
		log.Panicf("error reading the calibration document, %s\n", err.Error())
	}

	sum := 0
	for _, l := range lines {
		if s, err := calculateCalbiration(l); err == nil {
			sum += s
		} else {
			log.Panicf("error reading the calibration for line {%s}, %s\n", l, err.Error())
		}
	}

	fmt.Println(sum)
}

func calculateCalbiration(line string) (int, error) {
	var firstDigit byte
	var lastDigit byte

	l := 0
	r := len(line) - 1
	for l < len(line) {
		x := line[l]
		if '0' <= x && x <= '9' {
			lastDigit = x
		}

		y := line[r]
		if '0' <= y && y <= '9' {
			firstDigit = y
		}

		l++
		r--
	}

	value := (int(firstDigit - '0') * 10) + int(lastDigit - '0')
	return value, nil
}

func readDocument(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("calibration document is corrupt or does not exist, %w", err)
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("calibration document contents are invalid or file is corrupt, %w", err)
	}

	return lines, nil
}


package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	reZero  = regexp.MustCompile(`zero`)
	reOne   = regexp.MustCompile(`one`)
	reTwo   = regexp.MustCompile(`two`)
	reThree = regexp.MustCompile(`three`)
	reFour  = regexp.MustCompile(`four`)
	reFive  = regexp.MustCompile(`five`)
	reSix   = regexp.MustCompile(`six`)
	reSeven = regexp.MustCompile(`seven`)
	reEight = regexp.MustCompile(`eight`)
	reNine  = regexp.MustCompile(`nine`)
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
		if s, err := calibrate([]byte(l)); err == nil {
			sum += s
		} else {
			log.Panicf("error reading the calibration for line {%s}, %s\n", l, err.Error())
		}
	}

	fmt.Println(sum)
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

func calibrate(line []byte) (int, error) {
	line = reZero.ReplaceAll(line, []byte("zero0zero"))
	line = reOne.ReplaceAll(line, []byte("one1one"))
	line = reTwo.ReplaceAll(line, []byte("two2two"))
	line = reThree.ReplaceAll(line, []byte("three3three"))
	line = reFour.ReplaceAll(line, []byte("four4four"))
	line = reFive.ReplaceAll(line, []byte("five5five"))
	line = reSix.ReplaceAll(line, []byte("six6six"))
	line = reSeven.ReplaceAll(line, []byte("seven7seven"))
	line = reEight.ReplaceAll(line, []byte("eight8eight"))
	line = reNine.ReplaceAll(line, []byte("nine9nine"))

	firstDigit := byte('a')
	lastDigit := byte('a')

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

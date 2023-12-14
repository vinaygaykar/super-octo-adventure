package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	reDigit         = regexp.MustCompile(`\d+`)
	reDigitOrSymbol = regexp.MustCompile(`[^\d.]+`)
)

func main() {
	if len(os.Args) < 1 {
		log.Panicln("invalid number of arguments passed; please provide path to engine schematic document")
	}

	path := os.Args[1]
	schematic, err := parseEngineSchematicDoc(path)
	if err != nil {
		log.Panicln("error while parsing engine schematic document")
	}

	partNumbers, err := getPartNumbers(schematic)
	if err != nil {
		log.Panicln("error while parsing part numbers")
	}

	sum := uint64(0)
	for _, n := range partNumbers {
		sum += uint64(n)
	}

	fmt.Println(sum)
}

func parseEngineSchematicDoc(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	schematics := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		schematics = append(schematics, []byte(scanner.Text()))
	}

	return schematics, nil
}

func getPartNumbers(schematic [][]byte) ([]int, error) {
	validParts := make([]int, 0)

	for rIdx, r := range schematic {

		pos := reDigit.FindAllIndex(r, -1)
		if pos == nil {
			continue
		}

		for _, p := range pos {
			if isLeftOrRightAdjacentToSymbol(r, p[0]-1) ||
				isLeftOrRightAdjacentToSymbol(r, p[1]) ||
				(rIdx > 0 && isTopOrBottomAdjacentToSymbol(r, schematic[rIdx-1], p)) ||
				(rIdx < len(schematic)-1 && isTopOrBottomAdjacentToSymbol(r, schematic[rIdx+1], p)) {
				num, err := strconv.Atoi(string(r[p[0]:p[1]]))
				if err != nil {
					return nil, err
				}
				validParts = append(validParts, num)
			}
		}
	}

	return validParts, nil
}

func isLeftOrRightAdjacentToSymbol(line []byte, pos int) bool {
	if pos < 0 || len(line) <= pos {
		return false
	}

	return reDigitOrSymbol.Match(line[pos : pos+1])
}

func isTopOrBottomAdjacentToSymbol(line, otherLine []byte, pos []int) bool {
	start := pos[0] - 1
	if start < 0 {
		start = 0
	}

	end := pos[1] + 1
	if end > len(otherLine) {
		end = len(otherLine)
	}

	return reDigitOrSymbol.Match(otherLine[start:end])
}

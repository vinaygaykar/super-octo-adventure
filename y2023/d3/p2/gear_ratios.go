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
	reDigitOrSymbol = regexp.MustCompile(`[^\d.]{1}`)

	reGearPart = regexp.MustCompile(`\*+`)
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

	sum := 0
	for _, n := range partNumbers {
		sum += n
	}
	fmt.Println(sum)

	gearRatios, err := getGearRatio(schematic)
	if err != nil {
		log.Panicln("error while parsing gear part numbers")
	}

	gearSum := 0
	for _, n := range gearRatios {
		gearSum += n
	}
	fmt.Println(gearSum)
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

func getGearRatio(schematic [][]byte) ([]int, error) {
	ratios := make([]int, 0)

	for rIdx, r := range schematic {
		pos := reGearPart.FindAllIndex(r, -1)

		for _, p := range pos {
			var topLine, bottomLine []byte
			if rIdx > 0 {
				topLine = schematic[rIdx-1]
			}

			if rIdx < len(schematic)-1 {
				bottomLine = schematic[rIdx+1]
			}

			nums, err := getAdjacentNumbers(p[0], r, topLine, bottomLine)
			if err != nil {
				log.Panicf("error while finding adjacent number to gear at pos %d\n", p[0])
			}
			if len(nums) == 2 {
				ratios = append(ratios, nums[0]*nums[1])
			}
		}
	}

	return ratios, nil
}

func getAdjacentNumbers(pos int, line, topLine, bottomLine []byte) ([]int, error) {
	res := make([]int, 0)
	for _, p := range reDigit.FindAllIndex(line, -1) {
		// check nums on left & right of `*`
		if p[1] == pos || p[0] == pos+1 {
			num, err := strconv.Atoi(string(line[p[0]:p[1]]))
			if err != nil {
				return nil, err
			}
			res = append(res, num)
		}
	}

	// check nums on top of `*`
	for _, p := range reDigit.FindAllIndex(topLine, -1) {
		if p[0]-1 <= pos && pos <= p[1] && len(res) < 3 {
			num, err := strconv.Atoi(string(topLine[p[0]:p[1]]))
			if err != nil {
				return nil, err
			}
			res = append(res, num)
		}
	}

	// check nums on bottom of `*`
	for _, p := range reDigit.FindAllIndex(bottomLine, -1) {
		if p[0]-1 <= pos && pos <= p[1] && len(res) < 3 {
			num, err := strconv.Atoi(string(bottomLine[p[0]:p[1]]))
			if err != nil {
				return nil, err
			}
			res = append(res, num)
		}
	}

	return res, nil
}
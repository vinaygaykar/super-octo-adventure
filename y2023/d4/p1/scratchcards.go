package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	reDigit = regexp.MustCompile(`\d+`)
)

type Card struct {
	id    int
	win   []int
	found []int
}

func main() {
	if len(os.Args) < 1 {
		log.Panicln("invalid number of arguments passed; please provide path to cards table")
	}

	path := os.Args[1]
	cards, err := parseCards(path)
	if err != nil {
		log.Panicln("error while parsing cards table")
	}

	sum := 0
	for _, card := range *cards {
		winCount := -1

		for _, f := range card.found {
			if slices.Contains(card.win, f) {
				winCount++
			}
		}

		sum += int(math.Pow(2, float64(winCount)))
	}
	fmt.Println(sum)
}

func parseCards(path string) (*[]Card, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		card, err := parseCard(scanner.Text())
		if err != nil {
			return nil, err
		}

		cards = append(cards, *card)
	}

	return &cards, nil
}

func parseCard(cardStr string) (*Card, error) {
	cardInfo := strings.Split(cardStr, ": ")
	id, err := strconv.Atoi(string(reDigit.Find([]byte(cardInfo[0]))))
	if err != nil {
		return nil, err
	}

	numsInfo := strings.Split(cardInfo[1], " | ")

	win := make([]int, 0)
	for _, winStr := range reDigit.FindAllString(numsInfo[0], -1) {
		w, err := strconv.Atoi(winStr)
		if err != nil {
			return nil, err
		}
		win = append(win, w)
	}

	found := make([]int, 0)
	for _, foundStr := range reDigit.FindAllString(numsInfo[1], -1) {
		f, err := strconv.Atoi(foundStr)
		if err != nil {
			return nil, err
		}

		found = append(found, f)
	}

	return &Card{
		id:    id,
		win:   win,
		found: found,
	}, nil
}

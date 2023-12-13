package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id      int
	subsets []CubeSet
}

func main() {
	if len(os.Args) < 1 {
		log.Panicln("invalid number of arguments passed, pass in the game information file path")
	}

	path := os.Args[1]
	games, err := parseAllGames(path)
	if err != nil {
		log.Panicf("error while reading game info, %s\n", err.Error())
	}

	sum := 0
	sumOfPower := 0
	for _, game := range games {
		if isGamePossible(game, 12, 13, 14) {
			sum += game.id
		}

		set := lowestPossibleSet(game)
		sumOfPower += (set.red * set.green * set.blue)
	}

	fmt.Println(sum)
	fmt.Println(sumOfPower)
}

func parseAllGames(path string) ([]*Game, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("calibration document is corrupt or does not exist, %w", err)
	}

	games := make([]*Game, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if game, err := parseGame(scanner.Text()); err == nil {
			games = append(games, game)
		} else {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("calibration document contents are invalid or file is corrupt, %w", err)
	}

	return games, nil
}

func parseGame(gameStr string) (*Game, error) {
	gameInfo := strings.Split(gameStr, ": ")
	id, err := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
	if err != nil {
		return nil, err
	}

	sets := make([]CubeSet, 0)
	for _, setStr := range strings.Split(gameInfo[1], "; ") {
		set := CubeSet{}
		for _, cubeStr := range strings.Split(setStr, ", ") {
			setInfo := strings.Split(cubeStr, " ")
			count, err := strconv.Atoi(setInfo[0])
			if err != nil {
				return nil, err
			}

			switch setInfo[1] {
			case "red":
				set.red = count
			case "blue":
				set.blue = count
			default:
				set.green = count
			}
		}
		sets = append(sets, set)
	}

	return &Game{id, sets}, nil
}

func isGamePossible(game *Game, maxRed, maxGreen, maxBlue int) bool {
	for _, set := range game.subsets {
		if set.red > maxRed || set.green > maxGreen || set.blue > maxBlue {
			return false
		}
	}

	return true
}

func lowestPossibleSet(game *Game) CubeSet {
	red := 0
	green := 0
	blue := 0

	for _, set := range game.subsets {
		if set.red > red {
			red = set.red
		}

		if set.green > green {
			green = set.green
		}

		if set.blue > blue {
			blue = set.blue
		}
	}

	return CubeSet{red, green, blue}
}

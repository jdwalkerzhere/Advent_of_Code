package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func getData(filepath string) ([]string, error) {
	filehandle, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(filehandle)
	data := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data, nil
}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func search(data []string, r, c int, word string) int {
	xmasCount := 0

	for _, direction := range directions {
		mvY := direction[0]
		mvX := direction[1]
		xmasCount += searchDirections(data, r, c, mvY, mvX, word)
	}
	return xmasCount
}

func searchDirections(data []string, r, c, mvY, mvX int, word string) int {
	mvR := r + mvY
	mvC := c + mvX
	if mvR < 0 || mvR >= len(data) || mvC < 0 || mvC >= len(data[r]) {
		return 0
	}
	// we can move in the new direction
	if data[mvR][mvC] == word[0] {
		if len(word) == 1 { // We've completed an XMAS Search
			return 1
		}
		// Go look for nextLetter
		return searchDirections(data, mvR, mvC, mvY, mvX, word[1:])
	}
	// The letter in the direction is wrong
	return 0
}

func partOne(data []string, word string) int {
	xmasCount := 0
	for r, row := range data {
		for c := range row {
			if data[r][c] == word[0] {
				xmasCount += search(data, r, c, word[1:])
			}
		}
	}
	return xmasCount
}

var dirsTwo = [][]int{
	{-1, -1, 1, 1},
	{-1, 1, 1, -1},
}

func searchDirsTwo(data []string, r, c, mvYp, mvXp, mvYo, mvXo int) int {
	mvRp := r + mvYp
	mvCp := c + mvXp
	mvRo := r + mvYo
	mvCo := c + mvXo
	if mvRp < 0 || mvRp >= len(data) || mvCp < 0 || mvCp >= len(data[r]) {
		return 0
	}
	if mvRo < 0 || mvRo >= len(data) || mvCo < 0 || mvCo >= len(data[r]) {
		return 0
	}
	leftSide := string(data[mvRp][mvCp])
	rightSide := string(data[mvRo][mvCo])
	if leftSide == "M" && rightSide == "S" || leftSide == "S" && rightSide == "M" {
		return 1
	}
	return 0
}

func searchTwo(data []string, r, c int) int {
	xmasCount := 0
	for _, direction := range dirsTwo {
		mvYp := direction[0]
		mvXp := direction[1]
		mvYo := direction[2]
		mvXo := direction[3]
		xmasCount += searchDirsTwo(data, r, c, mvYp, mvXp, mvYo, mvXo)
	}
	if xmasCount == 2 {
		return 1
	}
	return 0
}

func partTwo(data []string) int {
	xmasCount := 0
	for r, row := range data {
		for c := range row {
			if string(data[r][c]) == "A" {
				xmasCount += searchTwo(data, r, c)
			}
		}
	}
	return xmasCount
}

func main() {
	flag.Parse()
	args := flag.Args()
	filepath := args[len(args)-1]
	data, err := getData(filepath)

	if err != nil {
		log.Fatal(err)
	}

	searchOne := "XMAS"
	answerOne := partOne(data, searchOne)
	fmt.Printf("Answer One: %d\n", answerOne)
	answerTwo := partTwo(data)
	fmt.Printf("Answer Two: %d\n", answerTwo)
}

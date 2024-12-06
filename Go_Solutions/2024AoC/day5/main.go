package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("No file path argument given, please provide the file")
	}
	filepath := args[len(args)-1]
	fileContents, err := readFile(filepath)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	answerOne, answerTwo := parts(fileContents)
	fmt.Printf("Sum Middle Page of Correct Page Ordering: %d\n", answerOne)
	fmt.Printf("Sum Middle Page of Incorrect Ordering: %d\n", answerTwo)
}

func readFile(filepath string) ([]string, error) {
	filehandle, err := os.Open(filepath)

	if err != nil {
		return nil, fmt.Errorf("Error reading the file: %w\n", err)
	}

	scanner := bufio.NewScanner(filehandle)
	fileContents := []string{}

	for scanner.Scan() {
		fileContents = append(fileContents, scanner.Text())
	}

	return fileContents, nil
}

func parts(fileContents []string) (int, int) {
	rules, updates, err := parseRulesUpdates(fileContents)

	if err != nil {
		log.Fatal(err)
	}

	sumValidMiddle := 0
	sumInvalidMiddle := 0
	for _, update := range updates {
		if ok := validUpdate(rules, update); !ok {
			reordered := makeValid(rules, update)
			middle := len(reordered) / 2
			sumInvalidMiddle += reordered[middle]
			continue
		}
		middle := len(update) / 2
		sumValidMiddle += update[middle]
	}
	return sumValidMiddle, sumInvalidMiddle
}

func parseRulesUpdates(fileContents []string) (map[int][]int, [][]int, error) {
	rules := make(map[int][]int)
	updates := [][]int{}

	parseRules := true
	for _, line := range fileContents {
		if line == "" {
			parseRules = false
			continue
		}
		if parseRules {
			// split line at |
			numStrings := strings.Split(line, "|")
			lft, rgt := numStrings[0], numStrings[1]
			lhs, err := strconv.Atoi(lft)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing into number: %s | %w\n", lft, err)
			}
			rhs, err := strconv.Atoi(rgt)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing into number: %s | %w\n", rgt, err)
			}
			// Now we need to figure out the map
			if _, ok := rules[lhs]; !ok {
				rules[lhs] = []int{rhs}
				continue
			}
			rules[lhs] = append(rules[lhs], rhs)
			continue
		}
		// parsing Updates
		numSlice := []int{}
		for _, num := range strings.Split(line, ",") {
			integer, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing into number: %s | %w\n", num, err)
			}
			numSlice = append(numSlice, integer)
		}
		updates = append(updates, numSlice)
	}
	return rules, updates, nil
}

func validUpdate(rules map[int][]int, update []int) bool {
	seen := make(map[int]bool)
	for _, page := range update {
		if len(seen) == 0 {
			seen[page] = true
			continue
		}
		for _, s := range rules[page] {
			if _, ok := seen[s]; ok {
				return false
			}
		}
		seen[page] = true
	}
	return true
}

func makeValid(rules map[int][]int, update []int) []int {
	reordered := []int{}
	for _, page := range update {
		successors := rules[page]
		index := 0
		for _, rightPage := range reordered {
			if contains(successors, rightPage) {
				break
			}
			index++
		}
		reordered = insert(reordered, index, page)
	}
	return reordered
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func insert(slice []int, insert, value int) []int {
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice[:insert])
	newSlice[insert] = value
	copy(newSlice[insert+1:], slice[insert:])
	return newSlice
}

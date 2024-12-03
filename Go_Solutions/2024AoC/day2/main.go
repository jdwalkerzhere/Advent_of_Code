package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	largestDiff = 3
)

func getReports(filepath string) ([][]int, error) {
	filehandle, err := os.Open(filepath)

	if err != nil {
		return nil, errors.New("Error opening filepath")
	}
	defer filehandle.Close()

	scanner := bufio.NewScanner(filehandle)

	reports := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		strInts := strings.Split(line, " ")
		levels := []int{}
		for _, level := range strInts {
			number, err := strconv.Atoi(level)

			if err != nil {
				return nil, errors.New("Problem converting string to integer")
			}
			levels = append(levels, number)
		}
		reports = append(reports, levels)
	}
	return reports, nil
}

func isSafe(lhs, rhs int, increasing bool) bool {
	if increasing {
		return rhs > lhs && rhs-lhs <= largestDiff
	}
	return rhs < lhs && lhs-rhs <= largestDiff
}

func getModified(report []int) [][]int {
	modified := [][]int{}

	for i := range len(report) {
		leftReport := report[:i]
		rightRport := report[i+1:]
		innerReport := []int{}
		innerReport = append(innerReport, leftReport...)
		innerReport = append(innerReport, rightRport...)
		modified = append(modified, innerReport)
	}
	return modified
}

func safeModified(modified [][]int, tolerance int) bool {
	for _, report := range modified {
		if safeReport(report, tolerance) {
			return true
		}
	}
	return false
}

func safeReport(report []int, tolerance int) bool {
	increasing := report[0] < report[1]
	for i := range len(report) - 1 {
		if !isSafe(report[i], report[i+1], increasing) {
			if tolerance <= 0 {
				return false
			}
			tolerance--
			modified := getModified(report)
			return safeModified(modified, tolerance)
		}
	}
	return true
}

func parts(reports [][]int, tolerance int) int {
	safeLevels := 0
	for _, report := range reports {
		safe := safeReport(report, tolerance)
		if safe {
			safeLevels++
		}
	}
	return safeLevels
}

func main() {
	flag.Parse()
	args := flag.Args()
	filepath := args[len(args)-1]
	reports, err := getReports(filepath)

	if err != nil {
		log.Fatal(err)
	}

	answerPartOne := parts(reports, 0)
	fmt.Printf("Number of Safe Reports: %d\n", answerPartOne)
	answerPartTwo := parts(reports, 1)
	fmt.Printf("Number of Safe Reports w/Tolerance of 1: %d", answerPartTwo)
}

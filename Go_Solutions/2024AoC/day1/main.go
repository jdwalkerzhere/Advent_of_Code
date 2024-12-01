package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getFloats(filepath string) ([]float64, []float64) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftFloats := []float64{}
	rightFloats := []float64{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := strings.Fields(line)
		left, err := strconv.ParseFloat(lineNumbers[0], 64)

		if err != nil {
			log.Fatal(err)
		}
		leftFloats = append(leftFloats, left)

		right, err := strconv.ParseFloat(lineNumbers[1], 64)

		if err != nil {
			log.Fatal(err)
		}
		rightFloats = append(rightFloats, right)
	}
	sort.Float64s(leftFloats)
	sort.Float64s(rightFloats)
	return leftFloats, rightFloats
}

func zip[T any](left []T, right []T) [][]T {
	shortest := min(len(left), len(right))
	zipped := make([][]T, shortest)
	for i := range shortest {
		zipped[i] = []T{left[i], right[i]}
	}
	return zipped
}

func partOne(leftFloats []float64, rightFloats []float64) int {
	totalDifferences := 0
	for _, pair := range zip(leftFloats, rightFloats) {
		left := pair[0]
		right := pair[1]
		absolute := int(math.Abs(left - right))
		totalDifferences += absolute
	}
	return totalDifferences
}

func convFloatToIntSlice(floatSlice []float64) []int {
	intSlice := make([]int, len(floatSlice))
	for i, float := range floatSlice {
		intSlice[i] = int(float)
	}
	return intSlice
}

func partTwo(leftFloats []float64, rightFloats []float64) int {
	leftNums := convFloatToIntSlice(leftFloats)
	rightNums := convFloatToIntSlice(rightFloats)
	numMap := make(map[int]int)

	for _, left := range leftNums {
		numMap[left] = 0
	}

	for _, right := range rightNums {
		if _, ok := numMap[right]; ok {
			numMap[right]++
		}
	}

	similarityScore := 0
	for num, frequency := range numMap {
		similarityScore += num * frequency
	}
	return similarityScore
}

func main() {
	flag.Parse()
	args := flag.Args()
	filepath := args[len(args)-1]
	leftFloats, rightFloats := getFloats(filepath)
	answerPartOne := partOne(leftFloats, rightFloats)
	fmt.Printf("Total of differences: %d\n", answerPartOne)
	answerPartTwo := partTwo(leftFloats, rightFloats)
	fmt.Printf("Similarity Score: %d\n", answerPartTwo)
}

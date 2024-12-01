package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getNums(filepath string) ([]int, []int) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftFloats := []int{}
	rightFloats := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := strings.Fields(line)
		left, err := strconv.Atoi(lineNumbers[0])

		if err != nil {
			log.Fatal(err)
		}
		leftFloats = append(leftFloats, left)

		right, err := strconv.Atoi(lineNumbers[1])

		if err != nil {
			log.Fatal(err)
		}
		rightFloats = append(rightFloats, right)
	}
	sort.Ints(leftFloats)
	sort.Ints(rightFloats)
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

func partOne(leftNums []int, rightNums []int) int {
	totalDifferences := 0
	for _, pair := range zip(leftNums, rightNums) {
		difference := max(pair[0], pair[1]) - min(pair[0], pair[1])
		totalDifferences += difference
	}
	return totalDifferences
}

func partTwo(leftNums []int, rightNums []int) int {
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
	leftNums, rightNums := getNums(filepath)
	fmt.Printf("Total of differences: %d\n", partOne(leftNums, rightNums))
	fmt.Printf("Similarity Score: %d\n", partTwo(leftNums, rightNums))
}

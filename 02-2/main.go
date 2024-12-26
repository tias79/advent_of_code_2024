package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []uint) bool {
	var safe bool = true
	for j := 1; j < len(levels); j++ {
		var diff uint = levels[j] - levels[j-1]
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}
	if safe {
		return true
	}

	safe = true
	for j := 1; j < len(levels); j++ {
		var diff uint = levels[j-1] - levels[j]
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}
	return safe
}

func removeElement(array []uint, i int) []uint {
	ret := make([]uint, 0)
	ret = append(ret, array[:i]...)
	return append(ret, array[i+1:]...)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var levels [][]uint
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var lineNumbers []uint
		for _, field := range fields {
			num, err := strconv.ParseUint(field, 10, 64)
			if err != nil {
				fmt.Errorf("error parsing number: %v", err)
				return
			}
			lineNumbers = append(lineNumbers, uint(num))
		}
		levels = append(levels, lineNumbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var safeReports uint64 = 0
	for i := 0; i < len(levels); i++ {
		fmt.Println(levels[i])

		if isSafe(levels[i]) {
			safeReports += 1
			continue
		}

		for j := 0; j < len(levels[i]); j++ {
			var newSlice = removeElement(levels[i], j)
			fmt.Println(newSlice)
			if isSafe(newSlice) {
				safeReports += 1
				break
			}
		}
	}

	fmt.Println("Result:", safeReports)
}

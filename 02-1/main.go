package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		var safe bool = true
		for j := 1; j < len(levels[i]); j++ {
			var diff uint = levels[i][j] - levels[i][j-1]
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}
		if safe {
			safeReports += 1
			continue
		}

		safe = true
		for j := 1; j < len(levels[i]); j++ {
			var diff uint = levels[i][j-1] - levels[i][j]
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}
		if safe {
			safeReports += 1
			continue
		}
	}

	fmt.Println("Result:", safeReports)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortUint64Slice(s []uint64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var column1, column2 []uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		// Ensure we have two fields
		if len(fields) == 2 {
			if num1, err := strconv.ParseUint(fields[0], 10, 64); err == nil {
				column1 = append(column1, num1)
			}
			if num2, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
				column2 = append(column2, num2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	m := make(map[uint64]uint64)
	for i := 0; i < len(column2); i++ {
		v, ok := m[column2[i]]
		if !ok {
			m[column2[i]] = 1
		} else {
			m[column2[i]] = v + 1
		}
	}

	var sum uint64 = 0
	for i := 0; i < len(column1); i++ {
		v, ok := m[column1[i]]
		if ok {
			sum += v * column1[i]
		}
	}

	fmt.Println("Result:", sum)
}

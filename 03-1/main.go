package main

import (
	"bufio"
	"fmt"
	"os"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func parseMul(reader *bufio.Reader) (bool, error) {
	bufPeek, peekErr := reader.Peek(3)
	if peekErr != nil {
		return false, peekErr
	}

	if string(bufPeek) == "mul" {
		for j := 0; j < 3; j++ {
			_, error := reader.ReadByte()
			if error != nil {
				return false, error
			}
		}
		return true, nil
	} else {
		return false, nil
	}
}

func parseParameters(reader *bufio.Reader) (int, int, error) {
	bufPeek, peekErr := reader.Peek(9)
	if peekErr != nil {
		return -1, -1, peekErr
	}

	var first int = 0
	var second int = 0
	var i int = 0
	if bufPeek[i] != '(' {
		return 0, 0, nil
	}
	i++
	for j := 0; j < 3; j++ {
		if isDigit(bufPeek[i]) {
			first *= 10
			first += (int(bufPeek[i]) - int('0'))
			i++
		}
	}
	if bufPeek[i] != ',' {
		return 0, 0, nil
	}
	i++
	for j := 0; j < 3; j++ {
		if isDigit(bufPeek[i]) {
			second *= 10
			second += (int(bufPeek[i]) - int('0'))
			i++
		}
	}
	if bufPeek[i] != ')' {
		return 0, 0, nil
	}
	for j := 0; j < i; j++ {
		_, error := reader.ReadByte()
		if error != nil {
			return -1, -1, peekErr
		}
	}
	return first, second, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var sum int = 0
	for {
		mul, error := parseMul(reader)
		if error != nil {
			break
		}
		if mul {
			first, second, error := parseParameters(reader)
			if error != nil {
				break
			}

			sum += first * second
		} else {
			_, error := reader.ReadByte()
			if error != nil {
				break
			}
		}
	}

	fmt.Println("Result:", sum)
}

package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	NumberStr = "0123456789"
)

func RunPartOne() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		firstPassPointer := 0
		secondPassPointer := len(line) - 1
		foundFirst := false

		for ; firstPassPointer < len(line); firstPassPointer++ {
			char := line[firstPassPointer]
			if strings.Contains(NumberStr, string(char)) {
				foundFirst = true
				break
			}
		}

		if !foundFirst {
			panic("no number found")
		}

		for ; secondPassPointer >= firstPassPointer; secondPassPointer-- {
			char := line[secondPassPointer]
			if strings.Contains(NumberStr, string(char)) {
				break
			}
		}

		num, err := strconv.Atoi(string(line[firstPassPointer]) + string(line[secondPassPointer]))
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println(sum)
}

var numStrMap = map[string]int{
	"zero":  0,
	"0":     0,
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}

func parseNumber(s string, i int) int {
	part := s[i:min(len(s), i+5)]
	for k, v := range numStrMap {
		if strings.HasPrefix(part, k) {
			return v
		}
	}
	return -1
}

func RunPartTwo() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	var pointer int
	for scanner.Scan() {
		line := scanner.Text()
		pointer = 0

		var numSlice []int

		for pointer = 0; pointer < len(line); pointer++ {
			num := parseNumber(line, pointer)
			if num != -1 {
				numSlice = append(numSlice, num)
				break
			}
		}
		for pointer = len(line) - 1; pointer > 0; pointer-- {
			num := parseNumber(line, pointer)
			if num != -1 {
				numSlice = append(numSlice, num)
				break
			}
		}

		var numStr string
        if len(numSlice) == 0 {
            log.Fatal("Num slice len 0, line: " + line)
        }
		if len(numSlice) == 1 {
			numStr = fmt.Sprint(numSlice[0]) + fmt.Sprint(numSlice[0])
		} else {
			numStr = fmt.Sprint(numSlice[0]) + fmt.Sprint(numSlice[1])
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println(sum)
}

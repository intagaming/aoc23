package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LimitRed   = 12
	LimitGreen = 13
	LimitBlue  = 14
)

func RunPartOne() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		colonPos := strings.Index(line, ":")
		idStr := line[5:colonPos]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		rounds := strings.Split(line[colonPos+2:], "; ")
		var maxR, maxG, maxB int
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				parts := strings.Split(cube, " ")
				cubeCount, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				switch parts[1] {
				case "red":
					maxR = max(maxR, cubeCount)
				case "green":
					maxG = max(maxG, cubeCount)
				case "blue":
					maxB = max(maxB, cubeCount)
				}
			}
		}

		if maxR <= LimitRed && maxG <= LimitGreen && maxB <= LimitBlue {
			sum += id
		}
	}

	fmt.Println(sum)
}

func RunPartTwo() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		colonPos := strings.Index(line, ":")
		// idStr := line[5:colonPos]
		// id, err := strconv.Atoi(idStr)
		// if err != nil {
		// 	panic(err)
		// }

		rounds := strings.Split(line[colonPos+2:], "; ")
		var maxR, maxG, maxB int
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				parts := strings.Split(cube, " ")
				cubeCount, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				switch parts[1] {
				case "red":
					maxR = max(maxR, cubeCount)
				case "green":
					maxG = max(maxG, cubeCount)
				case "blue":
					maxB = max(maxB, cubeCount)
				}
			}
		}

		sum += maxR * maxG * maxB
	}

	fmt.Println(sum)
}

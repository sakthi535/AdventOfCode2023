package main

import (
	"fmt"
	"os"
	"strings"
)

/*

	0 3 - 467


	467..114..
	...*......
	..35..633.
	.......#..
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..

*/

var sum = 0

func getNumbers(row int, col int, graph *[][]int) int {
	var value = 0

	var left = col
	var right = col

	for left >= 0 && (*graph)[row][left] >= 0 {
		left--
	}

	for right < len((*graph)[row]) && (*graph)[row][right] >= 0 {
		right++
	}

	for i := left + 1; i < right; i++ {
		value = 10*value + (*graph)[row][i]
		(*graph)[row][i] = -2
	}

	fmt.Println(value)
	return value
}

func findNumbers(i int, j int, graph *[][]int) {

	var rowStart = (i > 0)
	var rowEnd = (i < len(*graph)-1)

	var colStart = (j > 0)
	var colEnd = (j < len((*graph)[0])-1)

	var ct = 0
	var val = 1

	if rowStart {
		if colStart {
			if (*graph)[i-1][j] >= 0 {
				ct++
				val *= getNumbers(i-1, j, graph)
			}

			if (*graph)[i-1][j-1] >= 0 {
				ct++
				val *= getNumbers(i-1, j-1, graph)
			}
		}

		if colEnd {
			if (*graph)[i-1][j+1] >= 0 {
				ct++
				val *= getNumbers(i-1, j+1, graph)
			}
		}

		if (*graph)[i-1][j] >= 0 {
			ct++
			val *= getNumbers(i-1, j, graph)
		}
	}

	if rowEnd {
		if colStart {
			if (*graph)[i+1][j-1] >= 0 {
				ct++
				val *= getNumbers(i+1, j-1, graph)
			}
		}

		if colEnd {
			if (*graph)[i+1][j+1] >= 0 {
				ct++
				val *= getNumbers(i+1, j+1, graph)
			}
		}

		if (*graph)[i+1][j] >= 0 {
			ct++
			val *= getNumbers(i+1, j, graph)
		}
	}

	if colStart {
		if (*graph)[i][j-1] >= 0 {
			ct++
			val *= getNumbers(i, j-1, graph)
		}
	}

	if colEnd {
		if (*graph)[i][j+1] >= 0 {
			ct++
			val *= getNumbers(i, j+1, graph)
		}
	}

	if ct == 2 {
		sum += val
	}
	/*
		i-1, i, i+1
		j-1, j, j+1

	*/
}

func main() {

	data, _ := os.ReadFile("part2.txt")

	var lines = strings.Split(string(data), "\n")
	// fmt.Println(len(lines))

	var graph = make([][]int, len(lines)-1)

	for idx, line := range lines[:len(lines)-1] {
		// fmt.Println(line)
		graph[idx] = make([]int, len(line)-1)

		for j, c := range line[:len(line)-1] {

			var val = 0
			if c <= '9' && c >= '0' {
				val = int(c) - 48
			} else if c != '*' {
				val = -2
			} else {
				val = -1
			}

			graph[idx][j] = val
		}
	}

	for i, row := range graph {
		for j, col := range row {
			if col == -1 {
				findNumbers(i, j, &graph)
			}
		}
	}

	fmt.Println(sum)

}

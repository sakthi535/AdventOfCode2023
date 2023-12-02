package main

import (
	"fmt"
	"os"
	"strings"
)

func atoi(input string) (int, string) {
	var n int
	for i, b := range []byte(input) {
		b -= '0'
		if b > 9 {
			return n, input[i:]
		}
		n = n*10 + int(b)
	}
	return n, ""
}

func main() {
	data, _ := os.ReadFile("input.txt")

	var sum = 0

	var lines = strings.Split(string(data), "\n")

	for _, line := range lines {
		// fmt.Println(line)

		var game_split = strings.Split(string(line), ":")
		// fmt.Println(game_split[1])

		var gameId, _ = atoi(strings.Split(game_split[0], " ")[1])
		fmt.Println(gameId)

		var flag = true

		var games = strings.Split(game_split[1], ";")

		for _, round := range games {
			// fmt.Println(round)
			var cubes = (strings.Split(string(round), ","))

			for _, cube := range cubes {
				count := strings.Split(cube, " ")

				ct, _ := atoi(count[1])
				str := strings.TrimSpace(count[2])

				if str == "green" {
					if ct > 13 {
						flag = false
					}
				} else if str == "red" {
					if ct > 12 {
						flag = false
					}
				} else if str == "blue" {
					if ct > 14 {
						flag = false
					}
				}
				// fmt.Println(flag)

			}
		}

		fmt.Println(flag, gameId, gameId)
		if flag {
			sum += gameId
		}
	}

	fmt.Println(sum)

}

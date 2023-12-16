package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	moves = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{1, 1},
		{1, -1},
		{1, 0},
		{0, -1},
		{0, 1},
	}
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	cache := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cache = append(cache, []byte(line))
	}
	part2(cache)
}
func part1(cache [][]byte) {
	sum := 0
	for i := 0; i < len(cache); i++ {
		for j := 0; j < len(cache[i]); j++ {
			if cache[i][j] == '.' || (cache[i][j] >= '0' && cache[i][j] <= '9') {
				continue
			}

			for _, move := range moves {
				if isNumber(cache, i+move[0], j+move[1]) {
					sum += parseNumber(cache, i+move[0], j+move[1])
				}
			}
		}
	}
	fmt.Println(sum)
}
func part2(cache [][]byte) {
	sum := 0
	for i := 0; i < len(cache); i++ {
		for j := 0; j < len(cache[i]); j++ {
			if cache[i][j] != '*' {
				continue
			}
			nums := make([]int, 0)
			for _, move := range moves {
				if isNumber(cache, i+move[0], j+move[1]) {
					nums = append(nums, parseNumber(cache, i+move[0], j+move[1]))
				}
			}
			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}
	}
	fmt.Println(sum)
}
func isNumber(cache [][]byte, i, j int) bool {
	if i < 0 || i >= len(cache) {
		return false
	}
	if j < 0 || j >= len(cache[i]) {
		return false
	}
	if cache[i][j] >= '0' && cache[i][j] <= '9' {
		return true
	}
	return false
}
func parseNumber(cache [][]byte, i, j int) int {
	start := j
	end := j
	for start >= 0 && isNumber(cache, i, start-1) {
		start--
	}
	for len(cache[i]) > end && isNumber(cache, i, end+1) {
		end++
	}
	num := 0
	s := ""
	fmt.Println("=============================================")
	fmt.Println(start, end)
	for start <= end {
		curr := int(cache[i][start] - '0')
		s = s + string(cache[i][start])
		cache[i][start] = '.'
		num = (num * 10) + curr
		start++
	}
	fmt.Println(start, end)
	fmt.Println(s)
	fmt.Println(num)
	fmt.Println("=============================================")
	return num
}

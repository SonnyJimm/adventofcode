package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	part2(scanner)
}

func part1(scanner *bufio.Scanner) {
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ": ")
		data = strings.Split(strings.TrimSpace(data[1]), " | ")
		winnings := make(map[string]bool)
		for _, win := range strings.Split(data[0], " ") {
			win = strings.TrimSpace(win)
			if _, err := strconv.Atoi(win); err != nil {
				continue
			}
			winnings[win] = true
		}
		count := 0
		for _, num := range strings.Split(data[1], " ") {
			num = strings.TrimSpace(num)
			if winnings[num] {
				if count == 0 {
					count = 1
				} else {
					count <<= 1
				}
				// fmt.Println(num)
			}
		}
		fmt.Println(winnings)
		fmt.Println(count)
		sum += count
		// fmt.Println(winnings)
	}
	fmt.Println(sum)
}
func part2(scanner *bufio.Scanner) {
	// sum := 0
	wins := make(map[int]int)
	mem := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.TrimSpace(strings.Split(data[0], "Card ")[1]))
		data = strings.Split(strings.TrimSpace(data[1]), " | ")
		winnings := make(map[string]bool)
		for _, win := range strings.Split(data[0], " ") {
			win = strings.TrimSpace(win)
			if _, err := strconv.Atoi(win); err != nil {
				continue
			}
			winnings[win] = true
		}
		count := 0
		for _, num := range strings.Split(data[1], " ") {
			num = strings.TrimSpace(num)
			if winnings[num] {
				count++
				// fmt.Println(num)
			}
		}
		wins[id] = count
		mem[id] = -1
		// fmt.Println(winnings)
	}
	sum := 0
	for i := len(mem); i > 0; i-- {
		mem[i] = dp(mem, wins, i)
		sum += mem[i]
	}
	fmt.Println(sum)
}
func dp(mem, wins map[int]int, i int) int {
	if wins[i] == 0 {
		return 1
	}
	if mem[i] == -1 {
		copies := 0
		for index := wins[i]; index > 0; index-- {
			copies += dp(mem, wins, i+index)
		}
		mem[i] = copies + 1
	}

	return mem[i]
}

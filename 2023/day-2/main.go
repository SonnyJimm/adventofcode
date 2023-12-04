package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	defer file.Close()
	reader := bufio.NewScanner(file)
	result := 0
	for reader.Scan() {
		line := reader.Text()
		num := 0
		for i := 0; i < len(line); i++ {
			found := false
			if line[i] >= '0' && line[i] <= '9' {
				num += int(line[i] - '0')
				found = true
			}
			for index, val := range numbers {
				if strings.HasPrefix(line[i:], val) {
					num += index + 1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			found := false
			if line[i] >= '0' && line[i] <= '9' {
				num *= 10
				num += int(line[i] - '0')
				found = true
			}
			for index, val := range numbers {
				if strings.HasPrefix(line[i:], val) {
					num *= 10
					num += index + 1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		fmt.Println(num)
		fmt.Println(line)
		result += num
	}
	fmt.Println("calibrated result :", result)
}

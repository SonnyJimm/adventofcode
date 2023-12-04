package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	result := 0
	for reader.Scan() {
		line := reader.Text()
		num := 0
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				num += int(line[i] - '0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				num *= 10
				num += int(line[i] - '0')
				break
			}
		}
		fmt.Println(num)
		fmt.Println(line)
		result += num
	}
	fmt.Println("calibrated result :", result)
}

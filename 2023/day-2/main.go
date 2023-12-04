package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		// ok := true
		data := strings.Split(line, ": ")
		conf := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, val := range strings.Split(data[1], "; ") {
			for _, cube := range strings.Split(val, ", ") {
				detail := strings.Split(cube, " ")
				amount, _ := strconv.Atoi(detail[0])
				color := detail[1]
				if amount > conf[color] {
					conf[color] = amount
				}
			}
		}
		result += conf["red"] * conf["green"] * conf["blue"]
	}
	fmt.Println("By the configureation : ", result)
}

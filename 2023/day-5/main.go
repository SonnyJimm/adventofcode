package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	content := string(file)
	part1(content)
}

type mapsTo struct {
	to   string
	args [][3]int
}

func part1(content string) {
	inputs := strings.Split(content, "\n\n")
	mapsToTitle := make(map[string]*mapsTo)

	for i := 1; i < len(inputs); i++ {
		datas := strings.Split(inputs[i], "\n")
		title := strings.Split(strings.Split(datas[0], " ")[0], "-to-")
		curr := &mapsTo{to: title[1], args: make([][3]int, len(datas)-1)}
		mapsToTitle[title[0]] = curr
		for j := 1; j < len(datas); j++ {
			arg := strings.Split(datas[j], " ")
			curr.args[j-1][0], _ = strconv.Atoi(arg[0])
			curr.args[j-1][1], _ = strconv.Atoi(arg[1])
			curr.args[j-1][2], _ = strconv.Atoi(arg[2])
		}
		sort.Slice(curr.args, func(i, j int) bool { return curr.args[i][1] < curr.args[j][1] })
	}
	seeds := strings.Split(strings.Split(inputs[0], ": ")[1], " ")
	res := make([][]int, 0)
	for i := 0; i < len(seeds); i = i + 2 {
		if seeds[i] == "" {
			log.Fatal("PANIC")
		}
		if seeds[i+1] == "" {
			log.Fatal("PANIC i+1")
		}
		num, _ := strconv.Atoi(seeds[i])
		seedRange, _ := strconv.Atoi(seeds[i+1])
		res = append(res, getLocation(mapsToTitle, "seed", num, seedRange)...)
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	fmt.Println(res)
}

func getLocation(maps map[string]*mapsTo, curr string, num int, seedRange int) [][]int {
	mappedRanges := search(maps[curr].args, num, seedRange)
	if curr == "humidity" {
		return mappedRanges
	}
	mapsTo := make([][]int, 0)
	for _, mappedRange := range mappedRanges {
		mapsTo = append(mapsTo, getLocation(maps, maps[curr].to, mappedRange[0], mappedRange[1])...)
	}
	return mapsTo
}
func search(args [][3]int, num int, seedRange int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(args); i++ {
		dif := num - args[i][1]
		til := dif + seedRange
		if dif < 0 && til > 0 {
			res = append(res, []int{num, args[i][1] - num})
			dif = 0
			num = args[i][1]
			seedRange = (num + seedRange) - args[i][1]
			til = seedRange
		}
		if dif < 0 && til <= 0 {
			res = append(res, []int{num, seedRange})
			break
		}
		if dif >= 0 {
			if dif > args[i][2] {
				continue
			}
			if num+seedRange < args[i][1]+args[i][2] {
				res = append(res, []int{args[i][0] + dif, seedRange})
				break
			}
			res = append(res, []int{args[i][0] + dif, args[i][2] - dif})
			seedRange = (num + seedRange) - (args[i][1] + args[i][2])
			num = args[i][1] + args[i][2]
		}
	}
	if args[len(args)-1][2]+args[len(args)-1][1] <= num {
		res = append(res, []int{num, seedRange})
	}
	return res
}

package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"github.com/axhav/aoc2024/utils"
)

func star2() {
	sum := 0
	var aList,bList []int
	for _,line := range strings.Split(utils.ReadInput(), "\n") {
		if line == "" {
			continue
		}
		vals := strings.Split(line, "   ")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		aList = append(aList,a)
		b, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		bList = append(bList,b)
	}
	for _,v := range aList {
		c := utils.Count(bList,v)
		sum += c * v
	}
	fmt.Println(sum)
}

func star1() {
	sum := 0
	var aList,bList []int
	for _,line := range strings.Split(utils.ReadInput(), "\n") {
		if line == "" {
			continue
		}
		vals := strings.Split(line, "   ")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		aList = append(aList,a)
		b, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		bList = append(bList,b)
	}
	sort.Ints(aList)
	sort.Ints(bList)
	for i := range aList {
		sum += utils.Abs(aList[i] - bList[i])
	}
	fmt.Println(sum)
}


func main() {
	//star1()
	star2()
}

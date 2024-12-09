package main

import (
	"github.com/axhav/aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)


func star1() {
	sum := 0
	for _,line := range strings.Split(utils.ReadInput(), "\n") {
		elements := strings.Split(line, " ")
		item1,_ := strconv.Atoi(elements[0])
		ok := true
		for _,e := range elements[1:] {
			item2,_ := strconv.Atoi(e)
			if utils.Abs(item1-item2) != 1 {
				ok = false
				break
			}
			item1 = item2
		}
		if ok {
			sum += 1
		}
	}
	fmt.Println(sum)
}

func main()  {
	star1()
}

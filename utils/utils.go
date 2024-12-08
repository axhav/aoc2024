package main

import (
	"os"
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0  {
		return -x
	}
	return x
}

func Count[T comparable](slice []T, v T) int {
    count := 0
    for _, s := range slice {
        if v == s {
            count++
        }
    }
    return count
}

func ReadInput() string {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(data[:])
}



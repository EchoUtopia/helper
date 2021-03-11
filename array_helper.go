package leetcode

import (
	"encoding/json"
	"fmt"
)

func ParseIntSlice(in string) []int {
	out := []int{}
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		panic(err)
	}
	return out
}

func ParseIntMatrix(in string) [][]int {
	out := [][]int{}
	if err := json.Unmarshal([]byte(in), &out); err != nil {
		panic(err)
	}
	return out
}

func PrintIntMatrix(m [][]int) {
	for _, v := range m {
		fmt.Println(v)
	}
}

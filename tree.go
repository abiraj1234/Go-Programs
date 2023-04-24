package main

import (
	"fmt"
	"sort"
)

func generatePattern(input []int) [][]int {
	n := len(input)
	pattern := make([][]int, n)
	pattern[0] = input
	// fmt.Println(pattern)
	for i := 1; i < n; i++ {
		pattern[i] = make([]int, n-i)
		for j := 0; j < n-i; j++ {
			pattern[i][j] = pattern[i-1][j] & pattern[i-1][j+1]
		}
	}
	return pattern
}

func printPattern(pattern [][]int) {
	sort.Slice(pattern, func(i, j int) bool {
		return len(pattern[i]) < len(pattern[j])
	})
	for row := 0; row < len(pattern); row++ {
		numSpaces := len(pattern) - row
		for i := 0; i < numSpaces; i++ {
			fmt.Print(" ")
		}
		for col := 0; col <= row; col++ {
			fmt.Print(pattern[row][col], " ")
		}
		fmt.Println()
	}
}

func main() {
	input := []int{1, 1, 1, 1, 1, 1}
	pattern := generatePattern(input)
	printPattern(pattern)
}

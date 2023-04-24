package main

import "fmt"

func getSubset(slice []string, start, end int) []string {
	return slice[start:end]
}

func main() {
	var daysOfweeks []string
	daysOfweeks = []string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}

	subset := getSubset(daysOfweeks[:], 1, 6)

	for i := 0; i < len(subset); i++ {
		fmt.Println(subset[i])
	}
}

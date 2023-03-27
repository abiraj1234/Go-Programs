package main

import (
	"fmt"
	"strings"
)

func trunccateSentence(s string, k int) string {
	return strings.Join(strings.Split(s, " ")[:k], " ")
}

func main() {

	num := ("Hello how are you Contestant")
	k := 4
	fmt.Println(trunccateSentence(num, k))

}

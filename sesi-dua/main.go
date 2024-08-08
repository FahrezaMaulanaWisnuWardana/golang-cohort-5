package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "selamat malam"
	split := strings.Split(kata, "")

	words := make(map[string]int)
	for _, val := range split {
		fmt.Println(val)
		words[strings.ToLower(val)]++
	}
	fmt.Println(words)

}

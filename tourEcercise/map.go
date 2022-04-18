package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}
	a := strings.Fields(s)
	// fmt.Printf("Splitted: %s\n", split[1])
	// for i := 0; i < len(split)-1; i++ {
	// 	dem := 1
	// 	for j := i + 1; j < len(split); j++ {
	// 		if split[i] == split[j] {
	// 			dem++
	// 		}
	// 	}
	// 	m[split[i]] = dem
	// }
	for _, v := range a {
		m[v]++
	}
	return m
}

func main() {
	s := "I ate a donut. Then I ate another donut."
	m := WordCount(s)
	fmt.Println(m)
	// wc.Test(WordCount)
}

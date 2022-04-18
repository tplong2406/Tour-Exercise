package main

import "fmt"

type Contact struct {
	name        string
	phoneNumber string
}

var m map[string]Contact

var l = map[int]Contact{
	1: {"Tran phi long", "0933720316"},
	2: {"Hai", "0989777253"},
	3: {"Ba", "0937724325"},
}

var n = map[int]int{
	1: 1,
	2: 2,
	3: 3,
	4: 4,
}

func main() {
	// Create a tic-tac-toe board.
	m = make(map[string]Contact)
	m["Me"] = Contact{"Tran Phi Long", "0933720316"}
	fmt.Println(m["Me"])
	l[4] = Contact{"Bo", "0364153514"}
	fmt.Println(l)
	delete(l, 3)
	fmt.Println(l)
	v, ok := l[3]
	fmt.Println("value: ", v, "ok?", ok)
	fmt.Println(n)
}

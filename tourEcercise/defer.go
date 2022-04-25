package main

func main() {
	defer println("Count Suceed")
	for i := 0; i < 10; i++ {
		defer print(i, " ")
	}
}

package main

import "fmt"

func main() {
	fmt.Println(greeting())
	fmt.Println(farewell())
	fmt.Println(ponder())
}

func greeting() string {
	return "hello"
}

func farewell() string {
	return "goodbye"
}

func ponder() string {
	return "hmm"
}

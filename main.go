package main

import "fmt"

func main() {
	fmt.Println(greeting())
	fmt.Println(farewell())
}

func greeting() string {
	return "hello"
}

func farewell() string {
	return "goodbye"
}

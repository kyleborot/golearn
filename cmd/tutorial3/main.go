package main

import "fmt"

func main() {
	var printValue = "hello"
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

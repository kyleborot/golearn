package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	var myMap = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap)

	for name := range myMap {
		fmt.Printf("Name: %v\n", name)
	}

}

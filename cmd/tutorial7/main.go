package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value i is: %v\n", i)
	p = &i  //p refers to the MEMORY ADDRESS of i, they now reference the same int32 value in memory
	*p = 10 //set the value at the POINTED TO memory location to 10
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value i is: %v\n", i)

	var k int32 = 2
	i = k //this is different. It copies the value of k into i's memory location rather than USING the same memory location as i

	//this is useful for slices

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	//both slice and sliceCopy will change because underneath, slices contain pointers to an underlying array
	//we're copying pointers here

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is %v", result)
}

func square(thing2 [5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

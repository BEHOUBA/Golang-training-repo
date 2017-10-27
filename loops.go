package main

import "fmt"

func main() {
	x := 39
	var pt *int
	pt = &x
	//var pt2 **int

	///pt2 = &pt

	fmt.Printf("Value of pointer initialized is: %d", *pt)
}

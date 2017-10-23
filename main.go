package main

import (
	"fmt"
	"math"
)

func calCircleArea(radius float32) float32 {
	return radius * radius * math.Pi
}

func main() {
	fmt.Println(calCircleArea(50))
}

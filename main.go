package main

import (
	"fmt"
)

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknow", "Unknow"
	}
	return region, continent
}

func main() {
	fmt.Println(location("New York"))
}

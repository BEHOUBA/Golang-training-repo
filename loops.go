package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {

	//var manasse = person{"Manasse Kouame", 24, "LE HAVRE"}
	var manasse person
	manasse.name = "Kouame Behouba Manasse"
	manasse.age = 24
	manasse.city = "Le havre Normandie"
	fmt.Printf("this person name is: %s\n", manasse.name)
	fmt.Printf("this person age is: %d\n", manasse.age)
	fmt.Printf("this person live in: %s", manasse.city)

}

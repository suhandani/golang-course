package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var name2 = "Eko Kurniawan"
	fmt.Println(name2)

	name2 = "Eko Khannedy"
	fmt.Println(name2)

	name3 := "Eko Kurniawan"
	fmt.Println(name3)

	name3 = "Eko Khannedy"
	fmt.Println(name3)

	var (
		firstName = "Eko Kurniawan"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

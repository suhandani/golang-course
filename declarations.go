package main

import "fmt"

func main() {
	type NoKTP struct {
		name string
	}

	var p NoKTP
	p.name = "121212121212"

	fmt.Println(p)
}

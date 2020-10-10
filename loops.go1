package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 3 {
			continue
		}
	}

	slice := []int{1, 2, 3}

	println()
	println("For loop collection, longer syntax")
	for index := 0; index < len(slice); index++ {
		println(slice[index])
	}

	println()
	println("For loop collection, compact syntax")

	for j, v := range slice {
		println(j, v)
	}

	println()
	println("For loop collection, compact syntax for maps")

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for j, v := range wellKnownPorts {
		println(j, v)
	}

	println()
	println("For loop collection, compact syntax for maps, ignoring values")

	for k := range wellKnownPorts {
		println(k)
	}

	println()
	println("For loop collection, compact syntax for maps, ignoring keys")
	for _, v := range wellKnownPorts {
		println(v)
	}
}

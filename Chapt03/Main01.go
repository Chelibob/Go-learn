package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"Илья": 20,
		"Олех": 20,
		"Петя": 20,
	}
	ages["Максим"] = 20
	ages["Alex"] = 24

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}

	age, exists := ages["Антон"]
	if !exists {
		fmt.Println("Антона нет в списке")
	} else {
		fmt.Printf("Антону %d лет\n", age)
	}

	var mar map[string]int
	mar = make(map[string]int)

	fmt.Println(mar)
	fmt.Println("mar == nil ->", mar == nil)
}

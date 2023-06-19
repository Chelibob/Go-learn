package main

import "fmt"

func main() {
	var s = make([]int, 4, 4)
	copy(s, []int{1, 2, 3, 4})

	fmt.Println(s)
}

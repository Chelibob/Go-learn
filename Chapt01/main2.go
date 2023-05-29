package main

import "fmt"

func zero(x *int) {
	*x = 5
}
func swap(x, y *int) {
	*x, *y = *y, *x
}
func main() {
	x := new(int)
	zero(x)
	fmt.Println(*x)

	y, z := 3, 5
	fmt.Printf("y = %d, z = %d\n", y, z)
	swap(&y, &z)
	fmt.Println("Значения y и z после функции swap")
	fmt.Printf("y = %d, z = %d\n", y, z)
}

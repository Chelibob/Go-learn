package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Println(len(slice))
	fmt.Println("Срез до отзеркаливания")
	fmt.Println(slice)

	slice = reverse(slice)
	fmt.Println("Срез после отзеркаливания")
	fmt.Println(slice)

	arrPow(5)
}

func arrPow(n int) {
	num := 2
	for i := 0; i < n; i++ {
		fmt.Println(num)
		num *= 2
	}
}

// Функция, которую написал сам
func reverse1(arr []int) []int {
	var result []int
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

// Более элегантное решение из интернета
func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

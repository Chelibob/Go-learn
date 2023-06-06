package main

import "fmt"

func main() {
	var toDoList [3]string
	toDoList[0] = "купить хлеб"
	toDoList[1] = "купить молоко"
	toDoList[2] = "купить пиво"
	fmt.Printf("Количество элементов в списке %d\n", len(toDoList))

	for _, item := range toDoList {
		fmt.Printf("%s\n", item)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

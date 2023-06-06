package main

import "fmt"

func main() {
	toDoList := [4]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"купить лыбу",
	}

	var tasks []string
	tasks = toDoList[1:4]

	for i := range tasks {
		fmt.Println(tasks[i])
	}

	toDoList[3] = "купить пиуко"
	fmt.Println("---- ПОСЛЕ ИЗМЕНЕНИЯ toDoList ----")

	for i := range tasks {
		fmt.Println(tasks[i])
	}

}

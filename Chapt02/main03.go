package main

import "fmt"

func main() {
	toDoList := []string{
		"купить хлеб",
		"купить молоко",
		"купить творог",
	}

	toDoList = append(toDoList, "купить пиука", "купить лыбки", "позаниматься в зале", "приготовить еду")

	fmt.Println("Длина списка: ", len(toDoList))
	fmt.Println("Емкость списка: ", cap(toDoList))
}

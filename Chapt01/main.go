package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	printCircleArea(2)
}

func printCircleArea(radius int) {
	circleArea, err := calcCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: S=πr2\n")

	fmt.Printf("Площадь круга: %f32 см. кв.\n\n", circleArea)
}

func calcCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return 0.0, errors.New("Радиус не может быть меньше или равен нулю")
	}
	return float32(radius) * float32(radius) * math.Pi, nil
}

package main

import (
	"fmt"
	"math"
)

func main() {
	//multiplication  - множення
	// area == пr2    /окружність
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("area is:", area)

	//integer division - цілочисельне ділення
	half := 1 / 2
	fmt.Println("half wich integer division", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("half float:", halfFloat)

	//squaring (raising to the power)- зведення в квадрат (піднесення до степеня)
	badThreeSqured := 3 ^ 2
	fmt.Println("badThreeSqured", badThreeSqured)
	goodThreeSquared := math.Pow(3.0, 7.0)
	fmt.Println("goodThreeSqured", goodThreeSquared)

	//modulus
	remainder := 50 % 3 // -залишок
	fmt.Println("remainder", remainder)

	// unary operators
	x := 3
	x++
	fmt.Println("x is", x)

	x--
	x--
	fmt.Println("x is now", x)

	// in go
	var y = x // y = x++
	y++
	fmt.Println("y is", y)

}

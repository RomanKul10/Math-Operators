package main

import (
	"errors"
	"fmt"
)

func main() {

	/* //relational and conditonal operators - відносні та умовні оператори

	second := 31
	minute := 1

	if minute < 59 && second+1 > 59 {
		minute++
	}
	fmt.Println(minute) */

	a := 12
	b := 3

	/* 	if b != 0 {
		c := divideTwoNumbers(a, b)

		if c == 2 {
			fmt.Println("We found a two")
		}
	} */
	/* if b != 0 && divideTwoNumbers(a, b) == 2 {
		fmt.Println("Found 2")
	} */

	// this case never work
	/* if divideTwoNumbers(a, b) == 2 && b != 0 {
		fmt.Println("Found 2")
	} */
	c, err := divideTwoNumbers(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2")
		}
	}

}

/* func divideTwoNumbers(x, y int) int {
return x / y */

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

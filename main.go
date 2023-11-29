package main

import "fmt"

func main() {

	/* //modulus
	// does one number divide EXACLY INTO ANOTHER?
	x := 12
	y := 3
	if x%y == 0 {
		fmt.Println(y, "divides execly into", x)
	} else {
		fmt.Println(y, "does not divide execly into", x)
	}

	thisMonth := 4
	fmt.Println("The next month after ", thisMonth, "is", thisMonth+1) */

	for m := 1; m <= 12; m++ {
		fmt.Println("The month after", m, "is", m%12+1)

	}

}

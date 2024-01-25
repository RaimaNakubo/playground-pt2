package main

import (
	"fmt"
	"math"
)

// if statement basic syntax
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// starting if with a short statement declaration
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {

	//for loop basic syntax
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("for loop")
	fmt.Println(sum)

	//the init and post statements are optional, like C's "while" loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	/* Infinite loop
	for {

	}
	*/

	fmt.Println("if statement")
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println("if with a short statement")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

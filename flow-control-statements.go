package main

import (
	"fmt"
	"math"
)

// Sqrt implements square root function
func Sqrt(x float64) float64 {
	z := x / 2
	buffer := float64(0)
	change := float64(1)

	fmt.Println("Calculating square root of", x)
	for i := 0; i < 10; i++ {
		buffer = z
		z -= (z*z - x) / (2 * z)
		change = buffer - z
		fmt.Println("Iteration", i, ":", z, "change:", change)

		if change == 0 {
			fmt.Printf("Calculation completed in %v iterations\n", i+1)
			break
		}
	}
	return z
}

/*
the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z²,
 to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method.
 It works well for many functions but especially well for square root
*/

func main() {
	fmt.Println("re-created square root is  ", Sqrt(122))
	fmt.Println("standard lib square root is", math.Sqrt(122))
}

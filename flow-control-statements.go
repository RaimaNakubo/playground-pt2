package main

import "fmt"

func main() {

	//basic syntax
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
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
}

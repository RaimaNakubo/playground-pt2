package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	//switch statement basic syntax
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//any other OS
		fmt.Printf("%s.\n", os)
	}
	fmt.Println()

	//switch evaluation order
	fmt.Println("Switch evaluation order:")
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println()

	//Switch with no condition
	fmt.Println("Switch with no condition:")
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening.")
	}
	fmt.Println()

	//deferred function call
	fmt.Println("Deferred statement:")

	/*
		defer fmt.Println("forgor")
		fmt.Println("I")

		fmt.Println()
	*/

	//stacking deferred calls
	fmt.Println("Counting:")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done!")
}

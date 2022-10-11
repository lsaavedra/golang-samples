package main

import "fmt"

func main() {
	intSum := func(x, y int) int { return x + y } // this is a closure
	fmt.Println("5 + 4 =", intSum(5, 4))

	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}
	changeVar()
	fmt.Println("samp1 =", samp1) // notice that change variable value

}

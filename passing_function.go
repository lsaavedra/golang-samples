package main

import "fmt"

var pl = fmt.Println

func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	useFunc(sumVals, 5, 8)

}

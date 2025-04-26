package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b

}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {

	sum, diff := SumAndDiff(10, 4)

	fmt.Println("sum:", sum)
	fmt.Println("diff:", diff)

	mul := Multiply(10, 4)
	fmt.Println("mul:", mul)

	div, remain := Divide(10, 4)
	fmt.Println("div:", div)
	fmt.Println("remain:", remain)

}

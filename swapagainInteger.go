package main

import "fmt"

func swapagain() {

	var num1 int = 30
	var num2 int = 80
	var num3 int = 5
	var num4 int = 0
	fmt.Println("Before Swapping Three number is :", num1, num2, num3)

	num4 = num1
	num1 = num2
	num2 = num3
	num3 = num4

	fmt.Println("After Swapping Now this Three Number Is", num1, num2, num3)
}

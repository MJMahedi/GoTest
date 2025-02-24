package main

import "fmt"

func swapInteger() {
	var num1 int = 5
	var num2 int = 9
	var num3 int = 0
	fmt.Println("Number Before Swapping number1 & Number2 are :", num1, num2)

	num3 = num1
	num1 = num2
	num2 = num3

	fmt.Println("Number After Swapping :")
	fmt.Println("Num1 :", num1)
	fmt.Println("Num2 :", num2)

}

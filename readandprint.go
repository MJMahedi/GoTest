package main

import "fmt"

func readAndPrint() {
	fmt.Println("Now We Learn Read and Print")
	var input int = 0

	fmt.Print("Enter The Value :")
	fmt.Scanf("%d", &input)

	fmt.Println("You Entered :", input)
}

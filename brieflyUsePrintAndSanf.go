package main

import "fmt"

func briefly() {
	//Declare 4 integer type variables.
	var p int
	var r int
	var t int
	var si int

	fmt.Print("Enter principal: ")
	fmt.Scanf("%d", &p)

	fmt.Print("Enter rate: ")
	fmt.Scanf("%d", &r)

	fmt.Print("Enter time: ")
	fmt.Scanf("%d", &t)

	si = (p * r * t) / 100

	fmt.Println("Simple interest is: ", si)
}

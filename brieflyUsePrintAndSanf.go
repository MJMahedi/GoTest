package main

import "fmt"

func briefly() {
	//Declare 4 integer type variables.
	var p int
	var r int
	var t int
	var si int

	fmt.Print("Enter principal: ")
	fmt.Scanln(&p)

	fmt.Print("Enter rate: ")
	fmt.Scanln(&r)

	fmt.Print("Enter time: ")
	fmt.Scanln(&t)

	si = (p * r * t) / 100

	fmt.Println("Simple interest is: ", si)
}

package main

import "fmt"

func basicType() {
	fmt.Print("Basic Type Using In Go\n")
	//Boolean
	var b bool = true
	fmt.Printf("Boolean: %v, Type: %T\n", b, b)
	//Number
	//Intergers

	var i int = 42
	var i8 int8 = 127
	var i16 int16 = 32767

	fmt.Printf("Integer value: %v Type: %T\n", i, i)
	fmt.Printf("Integer value: %v Type: %T\n", i8, i8)
	fmt.Printf("Integer value: %v Type: %T\n", i16, i16)

}

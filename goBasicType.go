package main

import "fmt"

func basicType() {
	fmt.Print("Basic Type Using In Go\n")
	fmt.Print("----------Boolean--------\n")

	var b bool = true
	fmt.Printf("Boolean: %v, Type: %T\n", b, b)

	fmt.Printf("----Number & Intergers-----\n")
	var i int = 42
	var i8 int8 = 127
	var i16 int16 = 32767

	fmt.Printf("Integer value: %v Type: %T\n", i, i)
	fmt.Printf("Integer value: %v Type: %T\n", i8, i8)
	fmt.Printf("Integer value: %v Type: %T\n", i16, i16)

	fmt.Printf("----Floating Number-----\n")
	var f32 float32 = 3.14254571
	var f64 float64 = 3.142545718754

	fmt.Printf("Float32 Value: %v, Type %T\n", f32, f32)
	fmt.Printf("Float64 Value: %v, Type %T\n", f64, f64)

	fmt.Printf("----Complex Number-----\n")
	var c64 complex64 = 1 + 2i
	var c128 complex128 = complex(3, 4) // Another way to create complex number
	fmt.Printf("Complex64: %v, Type: %T\n", c64, c64)
	fmt.Printf("Complex128: %v, Type: %T\n", c128, c128)

	fmt.Printf("-------Strings-----\n")
	var mj string = "This is Strings, GO!"
	fmt.Printf("String Value: %v, Type: %T\n", mj, mj)

	fmt.Printf("---Runes (represents Unicode code Points)----\n")
	var r rune = 'A'
	fmt.Printf("Rune Value: %v, Type: %T\n", r, r)

	fmt.Printf("--Byte (Represnt ASCII Character)\n")
	var a byte = 'a'
	var newb byte = 'b'
	var newA byte = 'A'
	fmt.Printf("a Byte value: %v, Type: %T\n", a, a)
	fmt.Printf("b Byte value: %v, Type: %T\n", newb, newb)
	fmt.Printf("A Byte value: %v, Type: %T\n", newA, newA)
}

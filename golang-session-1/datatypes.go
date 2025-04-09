package main

import "fmt"

/*
	This is a Data Type Project.
	Testing a variation of int, float, bool, String.
	Understanding about nil and Zero Value
*/

// Combination of using int, float, bool, String
func main() {
	//int
	var firstInt = 76
	secondInt := -4

	fmt.Printf("Data type of firstInt: %T\n", firstInt)
	fmt.Printf("Data type of secondInt: %T\n\n", secondInt)

	//uint
	var UInt1 uint8 = 101
	var nonUInt2 int8 = -12

	fmt.Printf("Data type of uInt1: %T\n", UInt1)
	fmt.Printf("Data type of secondInt: %T\n\n", nonUInt2)

	//float
	var decimal1 float32 = 3.14

	fmt.Printf("Decimal default formatting: %f\n", decimal1)
	fmt.Printf("Decimal with nf formatting verb: %.2f\n\n", decimal1)

	//float
	var condition1 bool

	fmt.Printf("Is it permitted with Zero Value? %t\n\n", condition1)

	//String
	var message1 string = "Hello Tito\n\n"
	fmt.Print(message1)

	//String with backticks (``)
	var message2 = `Hello, welcome to Data Types training!
	it's me, Tito. Who is working on this project.`

	condition1 = true
	fmt.Println(message2)
	fmt.Printf("With result of Data types int: %T and %T, uint %T and Result float %.3f and with permission of the Head of our team is it permitted? %t", firstInt, secondInt, UInt1, decimal1, condition1)
}

/*
Result:
Data type of firstInt: int
Data type of secondInt: int

Data type of uInt1: uint8
Data type of secondInt: int8

Decimal default formatting: 3.140000
Decimal with nf formatting verb: 3.14

Is it permitted with Zero Value? false

Hello Tito

Hello, welcome to Data Types training!
it's me, Tito. Who is working on this project.
With result of Data types int: int and int, uint uint8 and Result float 3.140 and with permission of the Head of our team is it permitted? true
*/

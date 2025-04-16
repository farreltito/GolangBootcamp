package main

import (
	"fmt"
)

/*
	This is a Aliase project.
*/

func main() {
	//Aliase
	var typedataA uint8 = 11
	var typedataB byte //byte a.k.a type data of uint8

	typedataB = typedataA //no error, since byte == uint8
	_ = typedataB

	type second = uint

	var minute second = 60
	fmt.Printf("Minute type: %T\n", minute) // Minute type: uint
}

/*
	Note:
	- second is a.k.a uint
	- minute as second
	- minute tetap uint karena tipe datanya mengikuti dari second

	Result:
	Minute type: uint
*/

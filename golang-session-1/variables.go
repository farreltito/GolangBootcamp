package main

import "fmt" //fmt is a "Format Package"

/*
	This is a Hello World Project.
	Testing a variation of Println and Printf.
*/

// Variable without data type
func main() {
	var name string = "Tito" //set data
	var age int = 25

	name = "Mr 4k" //set data #2 or re-assign (Change)
	age = 20

	fmt.Printf("%T, %T", name, age) //%T sesuai dengan posisi, yg pertama buat name, yang kedua buat age
}

// Result: "This author name is%!(EXTRA string=Mr 4k)and his/her age is 20"

/*
	#### Note ####
	Advantages of Go languages:
	- All variables used.
*/

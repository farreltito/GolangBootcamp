package main

import "fmt" //fmt is a "Format Package"

/*
	This is a Variable Project.
	Testing a variation of Println and Printf.
*/

// Variable without Data Type
func main() {
	name := "Mr 4k" //set data
	age := 20

	fmt.Println(name, age)
	fmt.Printf("%T, %T", name, age)
}

/*
	Result:
	Mr 4k 20
	string, int
*/

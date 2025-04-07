package main

import "fmt"

/*
	This is a Variable Project.
	Testing a variation of Println and Printf.
*/

// Multiple Variable Declarations mix with Short Declaration
func main() {
	var companyname, companyage, companyAddress = "Sage", 2, "Jl. Jalan Happy"
	one, two, three := "1", 2.33, "3"

	fmt.Println(companyname, companyage, companyAddress)
	fmt.Println(one, two, three)
	fmt.Printf("%T, %T, %T", companyname, companyage, companyAddress)
	fmt.Printf("%T, %T, %T", one, two, three)
}

/*
	Result:
	Sage 2 Jl. Jalan Happy
	1 2.33 3
	string, int, stringstring, float64, string
*/

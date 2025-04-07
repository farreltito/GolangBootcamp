package main

/*
	This is a Variable Project.
	Testing a variation of Println and Printf.
*/

// Underscore Variable
func main() {
	var aVariable string
	var companyname, companyage, companyaddress = "Sage", 2, "Jl. Jalan Happy"

	// Underscore digunakan untuk menghindari error ketika variable tidak digunakan
	_, _, _, _ = aVariable, companyname, companyage, companyaddress
}

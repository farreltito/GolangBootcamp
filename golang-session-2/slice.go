package main

import (
	"fmt"
)

/*
	This is a Slice project.
*/

func main() {
	//Slice
	var agents = []string{"Sage", "Phoenix", "Sova", "Brimstone"}
	_ = agents

	//Slice (make function)
	var controllers = make([]string, 3)
	_ = controllers

	fmt.Printf("%#v\n", controllers) //will be empty strings

	//Slice (append function)
	//Assign or set data to controllers
	controllers[0] = "Brimstone"
	controllers[1] = "Viper"
	controllers[2] = "Omen"

	fmt.Printf("%#v\n", controllers)

	//this is where append process, from 3 become 6
	controllers = append(controllers, "Astra", "Harbor", "Clove")

	fmt.Printf("%#v\n", controllers)

	//Slice (append function with ellipsis)
	var duelists = []string{"Neon", "Jett", "Waylay"}
	var initiators = []string{"Sova", "Tejo", "Breach"}

	overrideData := copy(duelists, initiators)

	fmt.Println("Duelist:", duelists)
	fmt.Println("Initiator:", initiators)
	fmt.Println("Number copied of elements:", overrideData)

	//Slice (Slicing)

	//Slice (Combining slicing and append)

	//Slice (Backing array)

	//Slice (Cap function)

	//Slice (Creating a new backing array)
}

/*
	Note:

	Result:
*/

package main

import (
	"fmt"
)

/*
	This is a Array project.
*/

func main() {
	//Array
	//Method 1
	var noAgents [5]int
	noAgents = [5]int{1, 2, 3, 4, 5}
	//Method 2
	var agents = [4]string{"Sage", "Phoenix", "Sova", "Brimstone"}

	fmt.Printf("%#v\n", noAgents)
	fmt.Printf("%#v\n", agents)
	fmt.Println("")

	//Array (Modify Element Through Index)
	agents[1] = "Iso"
	agents[2] = "Astra"
	agents[3] = "Killjoy"

	fmt.Printf("%#v\n", agents)
	fmt.Println("")

	//Array (Loop through elements)
	//Method 1
	for i, j := range agents {
		fmt.Printf("No. Agents: %d, Agents: %s\n", i, j)
	}
	fmt.Println("")

	//Method 2
	for i := 0; i < len(agents); i++ {
		fmt.Printf("No. Agents: %d, Agents: %s\n", i, agents[i])
	}
	fmt.Println("")

	//Array (Multidimensional Array)
	//Kurung kotak awal row, kotak kedua column
	var oddNumbers [2][3]int
	oddNumbers = [2][3]int{{1, 3, 5}, {7, 9, 11}}

	for _, arr := range oddNumbers {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}

/*
	Note:
	- %#v adalah untuk memunculkan value
	- len stands for length. used for getting value of array

	Result:
	[5]int{1, 2, 3, 4, 5}
	[4]string{"Sage", "Phoenix", "Sova", "Brimstone"}

	[4]string{"Sage", "Iso", "Astra", "Killjoy"}

	No. Agents: 0, Agents: Sage
	No. Agents: 1, Agents: Iso
	No. Agents: 2, Agents: Astra
	No. Agents: 3, Agents: Killjoy

	No. Agents: 0, Agents: Sage
	No. Agents: 1, Agents: Iso
	No. Agents: 2, Agents: Astra
	No. Agents: 3, Agents: Killjoy

	1 3 5
	7 9 11
*/

package main

import "fmt"

/*
	This is a Variable Project.
	Testing a variation of Println and Printf.
*/

// Multiple Variable Declarations
func main() {
	var agent1, agent2, agent3 string = "Sage", "Brimstone", "Omen"

	var one, two, three int
	one, two, three = 1, 2, 3

	fmt.Println(agent1, agent2, agent3)
	fmt.Println(one, two, three)
	fmt.Printf("%T, %T, %T", agent1, agent2, agent3)
	fmt.Printf("%T, %T, %T", one, two, three)
}

/*
	Result:
	Sage Brimstone Omen
	1 2 3
	string, string, stringint, int, int
*/

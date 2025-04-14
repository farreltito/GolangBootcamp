package main

import (
	"fmt"
)

/*
	This is a Loopings project.
*/

func main() {
	//Loopings (first way of looping)
	//I want to make start from 1
	var firstLoop int = 1
	for firstLoop = 1; firstLoop <= 5; firstLoop++ {
		fmt.Println("Looping1:", firstLoop)
	}

	//Loopings (second way of looping)
	var secondLoop int = 0
	for secondLoop < 3 {
		fmt.Println("Looping2:", secondLoop)
		secondLoop++
	}

	//Loopings (third way of looping)
	thirdLoop := 0
	for {
		fmt.Println("Looping3:", thirdLoop)
		thirdLoop++

		if thirdLoop == 2 {
			break
		}
	}

	//Loopings (break and continue keywords)
	forthLoop := 1
	for forthLoop = 1; forthLoop <= 10; forthLoop++ {
		if forthLoop%2 == 1 {
			continue
		}

		if forthLoop > 9 {
			break
		}
		fmt.Println("Looping4:", forthLoop)
	}

	//Loopings (Nested Looping)
	for loopRow := 0; loopRow < 5; loopRow++ {
		for loopNumber := loopRow; loopNumber < 5; loopNumber++ {
			fmt.Print(loopNumber, " ")
		}

		fmt.Println()
	}

	//Loopings (Label)
outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Println("Loop on -", i+1)
		for j := 0; j < 5; j++ {
			if i == 4 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}

/*
	Result:
	Looping1: 1
	Looping1: 2
	Looping1: 3
	Looping1: 4
	Looping1: 5

	Looping2: 0
	Looping2: 1
	Looping2: 2
	Looping3: 0
	Looping3: 1

	Looping4: 2
	Looping4: 4
	Looping4: 6
	Looping4: 8

	0 1 2 3 4
	1 2 3 4
	2 3 4
	3 4
	4

	Loop on - 1
	0 1 2 3 4
	Loop on - 2
	0 1 2 3 4
	Loop on - 3
	0 1 2 3 4
	Loop on - 4
	0 1 2 3 4
	Loop on - 5
*/

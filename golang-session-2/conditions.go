package main

import (
	"fmt"
	"math/rand/v2"
)

/*
	This is a Conditions project.
*/

func main() {
	//Conditions (Temporary Variable)
	var thisYear = 2025

	if userAge := thisYear - 2000; userAge <= 18 {
		fmt.Println("You CAN'T enter this community. Because your age:", userAge, "\n")
	} else {
		fmt.Println("You CAN enter this community. Because your age:", userAge, "\n")
	}

	//Switch
	var max1 = 50
	var min1 = 0
	randomCombo50 := rand.IntN(max1-min1) + min1 //randomizer integer
	fmt.Println("Your combo:", randomCombo50)

	switch randomCombo50 {
	case 50:
		fmt.Println("Perfect Combo!\n")
	case 25:
		fmt.Println("Good Job!\n")
	default:
		fmt.Println("Try Again\n")
	}

	//Switch with Relational Operators
	var max2 = 100
	var min2 = 0
	randomScore100 := rand.IntN(max2-min2) + min2
	fmt.Println("Your score:", randomScore100)

	switch {
	case randomScore100 == 100:
		fmt.Println("Perfect score!\n")
	case (randomScore100 < 100) && (randomScore100 >= 51):
		fmt.Println("Study more!\n")
	default: //Braces for merapikan default
		{
			fmt.Println("Try Again and Study Harder!\n")
		}
	}

	//Switch (Fallthrough Keyword)
	var max3 = 12
	var min3 = 0
	randomNumber10 := rand.IntN(max3-min3) + min3
	fmt.Println("Your number:", randomNumber10)

	switch {
	case randomNumber10 == 10:
		fmt.Println("Perfect Number!\n")
	case (randomNumber10 < 10) && (randomNumber10 > 5):
		fmt.Println("Almost!")
		fallthrough
	case randomNumber10 < 7:
		fmt.Println("It's okay bad luck in randomizer, try again.\n")
	default: //jika kelebihan akan masuk ke default
		{
			fmt.Println("Whoops too far\n")
		}
	}

	//Nested Conditions
	var max4 = 10
	var min4 = 0
	randomNumber10v2 := rand.IntN(max4-min4) + min4
	fmt.Println("Your number (2):", randomNumber10v2)

	if randomNumber10v2 > 7 {
		switch randomNumber10v2 {
		case max4:
			fmt.Println("Perfecto!")
		default:
			fmt.Println("Naisu!")
		}
	} else {
		if randomNumber10v2 == 5 {
			fmt.Println("Not that bad!")
		} else if randomNumber10v2 == 3 {
			fmt.Println("Keep trying!")
		} else {
			fmt.Println("You can do it!")
			if randomNumber10v2 == 0 {
				fmt.Println("Try harder!")
			}
		}
	}

	/*
		Note:
		-

		Result:
		You CAN enter this community. Because your age: 25

		Your combo: 5
		Try Again

		Your score: 75
		Study more!

		Your number: 9
		Almost!
		It's okay bad luck in randomizer, try again.

		Your number (2): 1
		You can do it!
	*/
}

package main

import "fmt"

/*
	This is a Constant and Operators project.
*/

func main() {
	//Constant = should be set on first, and can't be changed
	const fullname string = "Farrel Tito"

	fmt.Printf("Hello %s\n", fullname)

	//Operators Arithmetic
	var result = 4 + 10*8      //pengerjaan sesuai matematika
	var result2 = (4 + 10) * 8 //pengerjaan sesuai kurung dahulu
	fmt.Println(result)
	fmt.Println(result2)

	//Operators Relational
	var condition1 bool = true             //true
	var condition2 bool = 4+5 > 10         //false
	var condition3 bool = 4 < 6            //true
	var condition4 bool = "tito" == "Tito" //false
	var condition5 bool = 2 != 2.22        //true
	var condition6 bool = 11 >= 12         //false

	fmt.Println("First Condition: ", condition1)
	fmt.Println("Second Condition: ", condition2)
	fmt.Println("Third Condition: ", condition3)
	fmt.Println("Fourth Condition: ", condition4)
	fmt.Println("Fifth Condition: ", condition5)
	fmt.Println("Sixth Condition: ", condition6)

	//Operators Logical
	const correct bool = true
	const notcorrect bool = false

	var logicalAnd = correct && notcorrect //AND
	fmt.Printf("Correct && NOT Correct \t(%t) \n", logicalAnd)

	var logicalOr = correct || notcorrect //OR
	fmt.Printf("Correct || NOT Correct \t(%t) \n", logicalOr)

	var logicalNot = !correct //NOT
	fmt.Printf("!Correct \t\t(%t) \n", logicalNot)

	/*
		Note:
		- \t itu tab dalam word

		Result:
		Hello Farrel Tito
		84
		112
		First Condition:  true
		Second Condition:  false
		Third Condition:  true
		Fourth Condition:  false
		Fifth Condition:  true
		Sixth Condition:  false
		Correct && NOT Correct  (false)
		Correct || NOT Correct  (true)
		!Correct                (false)
	*/
}

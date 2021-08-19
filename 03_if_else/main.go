package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		// if true
		println("Positive number")
	} else if num < 0 {
		println("Negative number")
	} else {
		// if not true
		println("You did not enter the number 5")
	}

	// ! operator
	// takes the opposite of the bool expression
	// !false => true "Not False"
	// !true => false "Not True"
	// num != 1 => "number not equal to 1"

	// if num != 10 {
	// 	println("Number is not 10")
	// }

	// if !false {
	// 	println("Not false evaluates to true")
	// }
	
	// || OR operator
	// If either expression evaluates to true
	// then it will execute the code block
	// if num < 0 || num > 10 {
	// 	println("Either negative or greater than 10")
	// }
	
	// && AND Operator
	// Both bool expressions must evaluate to true
	// for the code block to be executed
	// if num > 0 && num < 10 {
	// 	println("Between 1 and 9")
	// }

	// if num >= 5 {
	// 	println("Greater than or equal to 5")
	// }

	// if num < 5 {
	// 	println("Less than 5")
	// }

	// if num > 5 {
	// 	println("Greater than 5")
	// }

	// if num == 5 {
		//  if true
	// 	println("You entered the number 5")
	// }

	println("This will always run")

}
/**
	Summary:
	Use functions to perform basic operations and print some information to the terminal.

	Requirements:
	1. Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
	2. Write a function that returns any message, an call it from within fmt.Println()
	3. Write a function to add 3 numbers together, supplied as arguments, and return the answer
	4. Write a function that returns any number
	5. Write a function that returns any two numbers
	6. Add three numbers together using any combination of the existing functions. Print the result.
	* Call every function at least once
**/

package main

import "fmt"

func main() {
	// 1.
	greeting("Maxime")

	// 2.
	fmt.Println(say("Hello, world !"))

	// 3.
	total := add(1, 2, 3)

	// 4.
	sameTotal := giveMeBack(total)

	// 5.
	x, y := giveMeTwoNumbers()

	// 6.
	total = add(sameTotal, x, y)
	fmt.Println("The result is:", total)
}

func greeting(name string) {
	fmt.Println("Hello", name, "!")
}

func say(something string) string {
	return something
}

func add(a, b, c int) int {
	return a + b + c
}

func giveMeBack(number int) int {
	return number
}

func giveMeTwoNumbers() (int, int) {
	return 2, 8
}

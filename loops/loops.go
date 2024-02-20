/**
	Summary:
	Implement the class "FizzBuzz" problem using a `for` loop.

	Requirements:
	* Print integers 1 to 50, except:
	- Print "Fizz" if the integer is divisible by 3
	- Print "Buzz" if the integer is divisible by 5
	- Print "FizzBuzz" if the integer is divisible by both 3 and 5
**/

package main

import "fmt"

func main() {
	i := 1

	for i <= 50 {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

		i++
	}
}

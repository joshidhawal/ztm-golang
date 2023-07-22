//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {

	var favcolor string = "blue"
	var yearOfBirth, age = 1997, 26

	var (
		firstName = "Dhawal"
		lastName  = "Joshi"
	)

	ageInDays := age * 365

	fmt.Println("Favourite Colour =", favcolor)
	fmt.Println("Year Of Birth =", yearOfBirth, "age =", age)
	fmt.Println("firstName =", firstName, "lastName =", lastName)
	fmt.Println("Age in Days =", ageInDays)
}

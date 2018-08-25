/*
 *
 * The program demonstrates the use of pointers in go
 * By default parameters are passed by value, in the example below parameter will be passed by value
 *
 */
package main

import "fmt"

func main() {

	valueX := 0
	yPtr := new(int)

	// Pass the value and set its new value based on passed by reference
	fmt.Println("The memory address of valueX ", &valueX)
	//passByReference(&value)

	//Output the new value that should be increased by 10
	fmt.Println("The value of valueX prior to call by reference function ", valueX)
	passByReference(&valueX)
	fmt.Println("The value of valueX should be increased by 10 ", valueX)
	//Observe that the address of valueX has not changed
	fmt.Println("The memory address of value ", &valueX)
}

func passByReference(val *int) {

	// Print the memory address of val
	// Observe that the address of val changes on each call to this function
	// *val is the address where the value is stored for address of parameter passed into the function
	//
	fmt.Println("The address of val ", &val)

	*val += 10

}

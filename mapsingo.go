/*
*  This program demonstrates a number of key features of the go language
*
 */
package main

import "fmt"

//The env is configured to run linter automatically on save of the file/package
//

func main() {

	//defer operation - in this case run as the last function after the main function completes
	defer printFunc2() //this ismvery last function to be run before the program terminates
	printFunc1()       //this will be printed before func2 above - is the first func to be executed

	//Defer typically used as a clean-up function to be run as the last operation after everything else as
	//as been done

	//Run a factorial recursive function
	fmt.Println("The factorial of 10 = ", factorial(10))

	//Define a function in a function - a closure in go

	num := 3

	doubleNum := func() int {

		num *= 2

		return num
	}

	//Now call the closure multiple times
	fmt.Println("doubleNum once ", doubleNum())
	fmt.Println("doubleNum twice ", doubleNum())

	//Example using the map data structure - <key, value> mapping
	// map operation: create, add, delete, find, update, len -no. of eleemnts in the map

	// Create the map
	presidentMap := make(map[string]int)

	//Add items to the map
	presidentMap["Kennedy"] = 44
	presidentMap["Theodoore"] = 40
	presidentMap["Bush"] = 51

	//Print the length of the map
	fmt.Println("The len of the map is len")
	fmt.Println(len(presidentMap))
	fmt.Println("The len of the map is ", len(presidentMap))

	//Print elements from the map for a given key
	fmt.Println("The value of Kennedy")
	fmt.Println(presidentMap["Kennedy"])
	fmt.Println("Tha value of Kennedy is ", presidentMap["Kennedy"])
	//Delete items from the map for a given key
	delete(presidentMap, "Kennedy")
	fmt.Println(len(presidentMap)) //should now be two

	//find elements in the map for a given key

	//update the value of an element for a given key

	//Possible to define maps within maps - so the value for a given key is indeed a map which
	//itself defines a key value pair

	listofNums := []float64{1, 2, 3, 4, 5, 6}

	//Print the sum of the numbers
	fmt.Println(addThemUp(listofNums))

	//Print two return values
	num1, num2 := returnTwoValues(5) //func returns two values.
	fmt.Println(num1, num2)

	//Print the value from a variadic function - undefined number of arguments
	fmt.Println("The value of the variadic function ", subtractValues(1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7))
}
func printFunc2() {

	fmt.Println(" Print func2")
}

func printFunc1() {

	fmt.Println("Print func1")
}

//Function demonstrates passing an array of float, return a float
// use of the for loop with the range construct
//
func addThemUp(numbers []float64) float64 {

	sum := 0.0 //must initialised with a float

	for _, val := range numbers {
		// _ as the index is not required in the output
		sum += val

	}

	return sum
}

//Function returns two int values
//
func returnTwoValues(number int) (int, int) {

	return number + 1, number + 2
}

//Variadic function - undefined number of values input paramter
//
func subtractValues(args ...int) int {

	subValue := 0

	for _, value := range args {

		subValue -= value
	}

	return subValue
}

//Recursive function call
func factorial(num int) int {

	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}

}

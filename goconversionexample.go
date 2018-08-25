/*
This program demonstrates the use of the strconv library to convert between types e.g.
int to/from float
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Define a number of types float, int

	randInt := 5
	randFloat := 10.02030
	randStringInt := "100"
	randStringFloat := "200.01010"

	//Convert int to float
	fmt.Println(float64(randInt))

	//Convert float to int - Note the precision will be lost
	fmt.Println(int64(randFloat))

	//Convert string to int
	newInt, _ := strconv.ParseInt(randStringInt, 0, 64)
	fmt.Println(newInt)

	//Convert string to float
	newFloat, _ := strconv.ParseFloat(randStringFloat, 64)
	fmt.Println(newFloat)
}

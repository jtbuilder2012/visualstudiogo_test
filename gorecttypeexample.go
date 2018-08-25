/*
*
* This example demonstrates the use of the struct data type
*
 */

package main

import "fmt"

func main() {

	//Create an instance of the rectangle
	rect1 := Rectangle{length: 10, width: 10, top: 0, bottom: 0}

	//An alternative definition
	rect2 := Rectangle{10, 15, 0, 0}

	fmt.Println("The rectangle measures, length ", rect1.length, " and width ", rect1.width)
	fmt.Println("The rect2 width ", rect2.width)

	//Note the function below is attached to the Rectangle data structure
	fmt.Println("The area of rect1 is ", rect1.area())
	fmt.Println("The area of rect2 is ", rect2.area())
}

//Rectangle definiton  a struct data type....
type Rectangle struct {
	length float64
	width  float64
	top    float64
	bottom float64
}

func (rect *Rectangle) area() float64 {

	return rect.width * rect.length
}

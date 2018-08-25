package main

import (
	"fmt"
	"math"
)

func main() {

	//Create the shapes
	circle := Circle{radius: 10}
	circle1 := Circle{15}

	rect1 := Rectangle1{10, 20}
	rect2 := Rectangle1{length: 10, width: 30}

	//Calculate the area of each of the above shapes
	fmt.Println("The area of circle ", getArea(circle))
	fmt.Println("The area of circle1 ", getArea(circle1))
	fmt.Println("The area of the rectangle rect1", getArea(rect1))
	fmt.Println("The area of the rectangle rect1", getArea(rect2))
}

//The Shape interface definition
type Shape interface {
	area() float64
}

//The Rectangle1 definition
type Rectangle1 struct {
	length float64
	width  float64
}

//The Circle definition
type Circle struct {
	radius float64
}

func (rect Rectangle1) area() float64 {

	return rect.length * rect.width
}

func (circle Circle) area() float64 {

	return math.Pi * math.Pow(circle.radius, 2)
}

//This is the function that ties the interface area function to
//area of the specific shape
func getArea(shape Shape) float64 {
	//Maps the area to the struct type that is passed to this function
	return shape.area()
}

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	//var pi = math.Pi
	//var flag = false
	intarray := []int{4, 1, 5, 3, 5, 3}
	//Array to be copied to...
	var intarrayU [6]int
	//target to be copied to, source to be copied from
	copy(intarrayU[:], intarray)

	//var envArray = []string{}     //empty array
	//var envArray10 = [10]string{} //an array with 10 entry slots
	var myString = "This is hello world"

	//Run a number of strings operation
	fmt.Println(myString, "contains lo", strings.Contains(myString, "lo"))
	fmt.Println(myString, "index of lo", strings.Index(myString, "lo"))
	fmt.Println(myString, "count of ll", strings.Count(myString, "l"))
	fmt.Println(myString, "\"replace l with 3\"", strings.Replace(myString, "l", "3", -1))

	//Split a sting based on a given character, in this case a ','
	SampleCSVString := "1,2,3,5,6,7,8,9,"
	fmt.Println("Split the string", strings.Split(SampleCSVString, ","))

	log.Output(1, "Log output Error test line")
	strings.Contains(myString, "123")
	fmt.Println()
	sort.Ints(intarray)
	fmt.Println("Unsorted", intarrayU, "sorted", intarray)
	//strconv.
	//envArray := os.Environ()

	//Sort and print list of letters
	listOfLetters := []string{"C", "a", "1", "z", "Z", "p"}
	//sort the letters
	sort.Strings(listOfLetters)
	fmt.Println("Sorted list of letters ", listOfLetters)

	//Join a set of elements
	listOfNums := strings.Join([]string{"1", "2", "3"}, " ,")
	//Print the joined list of elements
	//Observe the output form of the list of elements
	fmt.Println(listOfNums)
}

func printEnvVariable() {

	fmt.Println()
	for _, e := range os.Environ() {
		//Split the env string of the form key=value
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

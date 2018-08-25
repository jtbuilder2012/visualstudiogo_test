package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//Create a file
	file, err := os.Create("Sample.txt")
	//Check if there was error creating the file
	if err != nil {
		log.Fatal(err)
	}

	//Write text to the file
	file.WriteString("This is the first write to the file")

	//Close the file
	file.Close()

	//Read text from the file
	stream, err := ioutil.ReadFile("Sample.txt")
	//Check if there was an error reading the file
	if err != nil {
		log.Fatal(err)
	}
	//Convert the stream to a string
	readString := string(stream)

	//Print the stream string
	fmt.Println(readString)

}

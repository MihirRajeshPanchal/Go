package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling")

	var text string

	fmt.Print("Enter some Text : ")
	fmt.Scanln(&text)

	file, err := os.Create("./myfile")

	checkNilError(err)

	length, err := io.WriteString(file, text)
	checkNilError(err)
	fmt.Println("Length is : ", length)

	readFile("./myfile")
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the File is : ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

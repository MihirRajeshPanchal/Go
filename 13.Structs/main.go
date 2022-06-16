package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Stu1 := Student{"Mihir", 18, 12, "IT"}
	fmt.Println("Details are : ", Stu1)
	fmt.Printf("Details are : %+v\n", Stu1)
	fmt.Printf("Name is : %+v and Field is : %+v  \n", Stu1.Name, Stu1.Field)

}

type Student struct {
	Name  string
	Age   int
	Class int
	Field string
}

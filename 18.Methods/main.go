package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Stu1 := Student{"Mihir", 18, 12, "IT"}
	// fmt.Println("Details are : ", Stu1)
	// fmt.Printf("Details are : %+v\n", Stu1)
	// fmt.Printf("Name is : %+v and Field is : %+v  \n", Stu1.Name, Stu1.Field)
	Stu1.GetName()
	Stu1.GetAge()
	Stu1.GetClass()
	Stu1.GetField()
}

type Student struct {
	Name  string
	Age   int
	Class int
	Field string
}

func (s Student) GetName() {
	fmt.Println("Name is : ", s.Name)
}
func (s Student) GetAge() {
	fmt.Println("Age is : ", s.Age)
}
func (s Student) GetClass() {
	fmt.Println("Class is : ", s.Class)
}
func (s Student) GetField() {
	fmt.Println("Field is : ", s.Field)
}

package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List : ", fruitList)
	fmt.Println("Length of Fruit List : ", len(fruitList))

	var vegList = [10]string{"Potato", "LadyFinger", "Cauliflower"}
	fmt.Println("Vegetable List is : ", vegList)
	fmt.Println("Length of Vegetable List is : ", len(vegList))
}

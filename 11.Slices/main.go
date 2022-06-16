package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Banana"}
	fmt.Printf("Type of Fruit List is : %T\n", fruitList)

	fmt.Println("List Before Append : ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("List After Append : ", fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println("List after deleting first node : ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 799
	highScores[1] = 588
	highScores[2] = 912
	highScores[3] = 643
	//highScores[4] = 234

	highScores = append(highScores, 803, 145, 344)

	fmt.Println("High Scores : ", highScores)
	fmt.Println("Boolean Value : ", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)

	fmt.Println("Sorted High Scores are : ", highScores)
	fmt.Println("Boolean Value : ", sort.IntsAreSorted(highScores))

	var courses = []string{"ReactJS", "Java", "Python"}
	fmt.Println("Courses : ", courses)

	var index int = 0
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after Delete : ", courses)

}

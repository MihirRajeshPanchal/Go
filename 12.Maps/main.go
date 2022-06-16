package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"

	fmt.Println("Languages are : ", languages)
	fmt.Println("PY stands for : ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("Languages are : ", languages)

	for key, value := range languages {
		fmt.Printf("Key : %v \t Value : %v \n", key, value)
	}
}

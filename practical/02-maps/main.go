package main

import "fmt"

func main() {
	grades := make(map[string]float32)
	grades["James"] = 42.39
	grades["Jes"] = 82.39
	grades["Serena"] = 82.41

	fmt.Println(grades)

	JamesGrade := grades["James"]

	fmt.Println(JamesGrade)

	delete(grades, "James")

	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}

package main

import (
	"fmt"
	"github.com/IshmaelKargbo/Learn-Go-in-One-Day/go-module/course"
	"strings"
)

func main() {
	fmt.Println("welcome to code 4 easy\n")
	var students []string

	for {
		name := course.Apply()
		students = append(students, name)
		fmt.Printf("Total student %v\n\n", len(students))

		for index, name := range students {
			names := strings.Fields(name)
			firstName := names[0]
			lastName := names[1]
			fmt.Printf("Index %v\n", index)
			fmt.Printf("First Name %v\n", firstName)
			fmt.Printf("Last Name %v\n\n", lastName)
		}
	}
}

package main

import "fmt"

func main() {
	var myArr [4]string

	myArr[1] = "test1"
	myArr[2] = "test2"
	myArr[3] = "test3"

	fmt.Println(myArr)
	fmt.Println(myArr[1])

	students := [5]string{"student1", "student2", "student3", "student4", "student5"}

	fmt.Println(students)

	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
		fmt.Printf("len=%d cap=%d %v\n", len(students), cap(students), students)
	}

}

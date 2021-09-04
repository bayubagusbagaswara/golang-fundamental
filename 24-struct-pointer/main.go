package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// Method dengan Pointer Receiver

func main() {

	student := Student{1, "Bayu Bagaswara", 3.23}

	fmt.Println(student.Name)

	graduate(&student)

	fmt.Println(student.Name)
}

func graduate(student *Student) {
	student.Name = student.Name + " S.T"
}

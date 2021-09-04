package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {

	student := Student{1, "Bayu Bagaswara", 3.23}

	fmt.Println(student.Name) // Bayu Bagaswara

	// panggil function graduate dan kirimkan parameter student yang berupa struct
	// pertama kita passing alamat memori dari struct student
	graduate(&student)

	fmt.Println(student.Name) // harapannya ini sudah diubah name nya

}

// kita ingin ngepassing struct ke dalam function
// tapi yang dipassing bukan value structnya, melaikan pointer alamat valuenya

// ini adalah function graduate yang menerima parameter struct Student
func graduate(student *Student) {
	// disini kita ingin mengubah nama studentnya, karena sudah lulus
	student.Name = student.Name + " S.T"

}

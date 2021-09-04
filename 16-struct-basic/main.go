package main

import "fmt"

// mengelompokan sebuah data dengan struktur yang kita inginkan lebih baik pake struct

type User struct {
	// kita definisikan propertinya
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {

	// kita bisa buat variabel untuk instance dari struct

	// Cara 1
	user := User{}
	// kita bisa set nilai di properties
	user.ID = 1
	user.FirstName = "Bayu"
	user.LastName = "Bagaswara"
	user.Email = "user@mail.com"
	user.IsActive = true
	fmt.Println(user)
	fmt.Println(user.FirstName)

	// Cara 2
	userTwo := User{ID: 2, FirstName: "Fuji", LastName: "Fauziah", Email: "fuji@mail.com", IsActive: true}
	fmt.Println(userTwo)

}

package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// membuat method untuk sebuah struct
// function display ini dimiliki atau instance dari object yang bertipe User
func (user User) display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func main() {

	// method adalah function yang dimiliki oleh sebuah object
	user1 := User{1, "Bayu", "Bagaswara", "bayu@mail.com", true}

	// memanggil function display
	result := user1.display()
	fmt.Println(result)

}

package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{1, "Bayu", "Bagaswara", "bayu@mail.com", true}

	fmt.Println(user)

	// karena struct itu adalah type, maka bisa kita perlakukan dia sama dengan type yang lain seperti int, string dll
	// dalam hal ini kita bisa perlakukan struct sebagai parameter di function
	displayUser := displayUser(user)

	fmt.Println(displayUser)

}

// function dengan parameter berupa struct

func displayUser(user User) string {
	// kita bisa akses properties dalam User
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}

package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	name        string
	admin       User
	users       []User
	IsAvailable bool
}

func main() {

	// sebuah Struct juga bisa menjadi atribut untuk Struct yang lainnya

	user1 := User{1, "Bayu", "Bagaswara", "bayu@mail.com", true}
	user2 := User{2, "Fuji", "Fauziah", "fuji@mail.com", true}

	displayUser(user1)

	users := []User{user1, user2}

	// Struct yang memilik properti dari struct User
	group := Group{"Gamer", user1, users, true}
	displayGroup(group)
}

func displayUser(user User) string {
	// kita bisa akses properties dalam User
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}
func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.users))
	fmt.Println("")

	for _, user := range group.users {
		fmt.Println("User: ", user.FirstName)
	}

}

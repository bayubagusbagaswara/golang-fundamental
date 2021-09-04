package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	name        string
	admin       User
	users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.users))
	fmt.Println("")

	for _, user := range group.users {
		fmt.Println("User: ", user.FirstName)
	}

}

func main() {

	// method adalah function yang dimiliki oleh sebuah object
	user1 := User{1, "Bayu", "Bagaswara", "bayu@mail.com", true}
	user2 := User{2, "Fuji", "Fauziah", "fuji@mail.com", true}

	// memanggil function display
	result := user1.Display()
	fmt.Println(result)

	users := []User{user1, user2}
	group := Group{"Gamer", user1, users, true}

	group.DisplayGroup()

}

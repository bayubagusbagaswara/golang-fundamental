package main

// karena User berada di package yang berbeda, maka kita import dulu
import (
	"21-exported-unexported-package/management"
	"fmt"
)

func main() {

	// kita bisa penggil User dengan menyebut nama package nya dahulu
	user1 := management.User{1, "Bayu", "Bagaswara", "bayu@mail.com", true}
	user2 := management.User{2, "Fuji", "Fauziah", "fuji@mail.com", true}

	// memanggil function display
	result := user1.Display()
	fmt.Println(result)

	Users := []management.User{user1, user2}
	group := management.Group{"Gamer", user1, Users, true}

	group.DisplayGroup()
}

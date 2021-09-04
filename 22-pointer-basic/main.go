package main

import "fmt"

func main() {

	// pointer adalah sebuah variabel yang isinya adalah suatu alamat di suatu nilai. Jadi yang disimpan itu bukan nilainya, tapi alamat memori dari nilai tersebut

	numberA := 5

	// menyimpan alamat dari numberA, dimana numberA menyimpan nilai 5
	numberB := &numberA

	fmt.Println(numberA) // 5
	fmt.Println(numberB) // alamat dari nilai 5

	// Dereferencing, agar kita bisa menerjemahkan nilai dari alamat yang ditunjuk oleh pointer
	fmt.Println(*numberB)

	// karena saling terikat, jika kita lakukan perubahan pada numberB, maka numberA juga akan berubah
	*numberB = 100

	fmt.Println(*numberB) // 100
	fmt.Println(numberA)  // menjadi 100
}

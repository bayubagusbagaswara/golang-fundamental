package main

import "fmt"

func main() {

	var numberA int = 5

	// *int artinya kita menyimpan alamat memori sebuah nilai yang bertipe integer
	var numberB *int = &numberA

	fmt.Println(numberA)  // 5
	fmt.Println(numberB)  // alamat memori
	fmt.Println(*numberB) // 5

	// Bagaimana kalau kita mengubah nilai di numberA?
	numberA = 20

	fmt.Println(numberA)  // 20
	fmt.Println(numberB)  // alamat memori sudah berubah
	fmt.Println(*numberB) // 20, ikut berubah

	// karena alamat numberA dan numberB itu sebenarnya sama (reference sama). Jadi, jika numberA diubah maka numberB juga ikut berubah

	number := 5
	fmt.Println("Nilai awal: ", number) // 5

	change(&number, 100)

	fmt.Println("Nilai akhir: ", number) // 100

}

func change(old *int, new int) {
	*old = new // ini adalah alamat memori
	fmt.Println("Di dalam function: ", *old)
}

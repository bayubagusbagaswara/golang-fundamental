package main

import "fmt"

func main() {

	// array adalah kumpulan dari variabel yang tipe datanya sama
	// jika kita ingin menampung banyak data di satu variable maka kita gunakan array

	// dalam array kita harus tentukan jumlah nilai didalam array itu sendiri, dan tipe data untuk arraynya

	// Array dengan jumlah data 5 bertipe data string. Jika hanya deklarasi array maka nilainya adalah default null (kosong)

	// var languages [5]string

	//  memasukkan nilai kedalam array dengan cara menambahkan index ke berapa

	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "Java"
	// languages[4] = "Python"

	// bisa juga langsung mengisi arraynya saat deklarasi awal
	// kalau penulisan horizontal tidak butuh koma diakhir
	// tapi penulisan vertikal membutuhkan koma diakhir
	// languages := [5]string{
	// 	"Ruby",
	// 	"Python",
	// 	"Javascript",
	// 	"Go",
	// 	"Java",
	// }

	// kita juga bisa membuat array tanpa harus menulis panjang jumlah data arraynya, tapi diganti dengan ...
	languages := [...]string{
		"Ruby",
		"Python",
		"Javascript",
		"Go",
		"Java",
		"C++",
		"Kotlin",
	}

	fmt.Println(languages)

	// mendapatkan panjang dari sebuah array
	length := len(languages)
	fmt.Println(length)

	// looping satu persatu untuk mengambil data di array
	for index, language := range languages {
		fmt.Println("Index", index, " Language: ", language)

	}
}

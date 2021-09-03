package main

import "fmt"

func main() {

	// for digunakan apabilan melakukan sesuatu secara berulang-ulang

	// misal mencetak kalimat sebanyak 100 kali

	// untuk i = 1, dengan syarat i adalah harus <= 100, maka lakukan perulangan
	for i := 1; i <= 100; i++ {
		fmt.Println("Saya sedang belajar Go : ", i)
	}

	// kenapa harus ada i++
	// saat pertama kali jalan maka nilai i = 1
	// jika i memenuhi syarat yakni i <= 100, maka jalankan perintah functionnya
	// setelah selesai, lalu nilai i akan diubah i + 1 (bertambah satu)
	// kemudian dicek ulang lagi i <= 100
	// dan seterusnya sampai syarat kondisinya bernilai false (tidak terpenuhi)
	// False saat i = 101, yang di print hanya sampai 100

	// Bentuk FOR Kedua, yakni While, tapi di GO tidak ada, hanya penulisannya adalah menggunakan for tapi dengan model while loop
	// juga harus dikasih batasnya, karena harapan kita suatu saat kondisinya bernilai FALSE (looping berhenti)
	j := 1
	for j <= 100 {
		fmt.Println("While Loop: ", j)
		j++
	}

	// Ketiga, for bisa digunakan untuk mengambil element dari sebuah collection
	title := "Golang the best language"

	// kita bisa mendapatkan tiap karakter
	// index adalah index tiap perulangan
	// data per huruf akan disimpan dalam letter (tapi yang disimpan adalah bytecodenya), kecuali dikonversi dulu
	for index, letter := range title {
		fmt.Println("index : ", index, " letter: ", string(letter))
	}

}

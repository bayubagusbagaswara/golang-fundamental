package main

import "fmt"

func main() {
	// memanggil function
	printMyResult()

	// memanggil function juga wajib memasukkan parameter/inputan
	helloWorld("Bayu")

	// memanggil function yang berfungsi melakukan penjumlahan, dan kita harus memasukkan inputan, dimana inputan tersebutlah yang akan diproses oleh function
	// Akan tetapi, kita harus buat variabel untuk menampung data balikan dari function tersebut atau bisa langsung diprint
	total := jumlah(5, 5)
	fmt.Println(total)
	// fmt.Println(jumlah(5, 5))

}

// contoh function
// parameter dan nilai balikan adalah opsional
func printMyResult() {
	fmt.Println("Saya sedang belajar Go")
}

// contoh function yang memiliki input/parameter
func helloWorld(name string) {
	fmt.Println("Hello World", name)

}

// contoh function yang mengembalikan sebuah nilai
// artinya function jumlah memiliki 2 parameter yakni a bertipe int dan b bertipe int, sedangkan balikan dari function ini juga bertipe int
func jumlah(a int, b int) int {
	return a + b
}

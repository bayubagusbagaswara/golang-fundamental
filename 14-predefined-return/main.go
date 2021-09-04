package main

import "fmt"

func main() {

	luas, keliling := calculate(20, 30)

	fmt.Println(luas)
	fmt.Println(keliling)

}

func calculate(panjang int, lebar int) (luas int, keliling int) {

	// karena sudah didefinisikan di function, maka kita cukup memakai variabel dengan nilai yang baru

	luas = panjang * lebar
	keliling = 2 * (panjang * lebar)

	// kita bisa hanya menulis return saja
	// secara otomatis dia akan mereturn variabel luas dan keliling, karena sudah didefinisikan di balikan functionnya
	// ini dinamakan predefined return, jadi mendefinisikan sebelum return

	return
}

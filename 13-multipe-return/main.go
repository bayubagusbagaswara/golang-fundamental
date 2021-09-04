package main

import "fmt"

func main() {

	// lalu kita buat variable untuk menampung masing-masing nilai luas dan keliling
	luas, keliling := calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)

	// bagaimana jika kita hanya menangkap satu nilai dari return function, maka harus kita ganti dengan tanda _ (ini namanya diignore)
	_, keliling1 := calculate(5, 5)
	fmt.Println(keliling1)

}

// balikannya bisa Luas dan Keliling
// bisa saja balikannya multiple, asal didefinisikan tipe data balikannya
// misal Luas adalah integer, dan Keliling adalah integer
func calculate(panjang int, lebar int) (int, int) {

	luas := panjang * lebar

	keliling := 2 * (panjang + lebar)

	return luas, keliling

}

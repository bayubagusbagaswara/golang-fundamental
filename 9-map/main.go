package main

import "fmt"

func main() {

	// deklarasi variable harus menggunakan keyword map
	// dan tentukan tipe data yang menjadi key dan value

	// artinya kita membuat variable myMap bertipe map dengan tipe data key nya String dan tipe data value nya String
	// var myMap map[string]string

	// pada kode program diatas variabel masih berupa deklarasi, sehingga nilai map masih nil (null)
	// kita perlu menginisialisasi variabel nya dengan nilai/data yang akan ditampungnya
	myMap := map[string]string{}

	// cara mengisi data ke map, harus menggunakan keynya
	myMap["ruby"] = "Ini Value Ruby"
	myMap["java"] = "Ini Value Java"

	fmt.Println(myMap)

	// untuk mendapatkan value dari map juga melalui key
	fmt.Println(myMap["java"])

	// melalukan iterasi untuk mengambil semua value di map
	for key, value := range myMap {
		fmt.Println("Key : ", key, " == Value : ", value)
	}

	// menghapus value di map
	delete(myMap, "ruby")

	fmt.Println(myMap)

	// mengecek suatu map apakah mengandung key tertentu
	fmt.Println(myMap["python"]) // hasilnya akan kosong, karena tidak ada key python

}

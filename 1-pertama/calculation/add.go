package calculation

// kita ingin function Add ini bisa dipanggil dari function main
// nama function harus diawali dengan Huruf Kapital jika ingin Public (bisa diakses package lain)
// jika diawali Huruf Kecil, maka package ini tidak bisa diakses dari package luar. Hanya bisa dipanggil dalam package yang sama (Private)

// function Add memiliki 2 parameter dan balikannya berupa integer
func Add(number int, numberTwo int) int {

	add := add(10, 10)              // 20
	return add + number + numberTwo // 20 + 5 + 5 = 30

}

func add(number int, numberTwo int) int {
	return number + numberTwo
}

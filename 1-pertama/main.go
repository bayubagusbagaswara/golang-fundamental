package main

import (
	"1-pertama/calculation"
	"fmt"
)

func main() {

	// didalam function main inilah source code akan dieksekusi

	fmt.Println("Hello, belajar Golang")

	// karena entity.go berada di pakcage main (packagenya sama), oleh itu bisa akses secara langsung disini
	// sentence := TestAja()

	// fmt.Println(sentence)

	// bisa memanggil file yang berbeda package
	// caranya dengan mengimport package yang ingin digunakan di main.go ini

	result := calculation.Add(5, 5)
	fmt.Println(result)

}

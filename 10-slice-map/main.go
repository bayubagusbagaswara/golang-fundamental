package main

import "fmt"

func main() {

	// map dan slice itu bisa digabung
	// untuk menampung banyak map kita bisa menggunakan slice

	// kita buat slice bertipe map, dimana map nya sendiri memiliki tipe data key string, dan value string
	students := []map[string]string{
		{
			"name":  "Bayu",
			"score": "A",
		},
		{
			"name":  "Bagus",
			"score": "B",
		},
		{
			"name":  "Bagaswara",
			"score": "C",
		},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}

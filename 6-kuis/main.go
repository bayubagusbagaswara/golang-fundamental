package main

import (
	"fmt"
)

func main() {

	kata := "Golang the best language"

	// cetak hanya index yang genap
	for index, huruf := range kata {
		if index%2 == 0 {
			fmt.Println("Index Genap: ", index, "Huruf : ", string(huruf))
		} else {
			continue
		}
	}

	// cetak hanya huruf vokal
	for index, huruf := range kata {

		hurufString := string(huruf)

		switch string(huruf) {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index: ", index, "Huruf: ", hurufString)

		}

		if hurufString == "a" || hurufString == "i" || hurufString == "u" || hurufString == "e" || hurufString == "o" {
			fmt.Println("Index: ", index, "Huruf: ", hurufString)
		}
	}
}

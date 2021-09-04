package main

import "fmt"

func main() {

	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var goodScores []int (slice)

	scoresB := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	// ambil tiap nilai dari array

	for index, score := range scoresB {
		// cek jika nilai >=90, maka masukkan ke slice
		if score >= 90 {
			fmt.Println("Index: ", index, " Score: ", score)
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)

	// hitung rata-rata
	scoresA := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int

	for _, score := range scoresA {
		// jumlahkan semua score
		total = total + score
	}
	fmt.Println("Total Nilai: ", total)
	average := float64(total) / float64(len(scoresA))
	fmt.Println("Rata-rata: ", average)
}

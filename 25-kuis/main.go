package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

// AddGame
// karena append di method ini hanya berlaku di dalam method ini saja
// sehingga kita perlu menambahkan pointer kedalam receiver
func (gamer *Gamer) AddGame(game string) {
	// setiap parameter yang masuk ke method AddGame akan disimpan kedalam struct Gamer
	gamer.Games = append(gamer.Games, game)

}
func main() {
	gamer := Gamer{Name: "Bayu"}
	gamer.AddGame("Winning Eleven")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}

}

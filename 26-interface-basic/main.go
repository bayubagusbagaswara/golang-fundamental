package main

// interface ini bisa kita samakan dengan Kontrak (aturan main yang harus dipenuhi)

type Luas interface {
	// kita buat sebuah kontrak
	// biasanya kontraknya berupa sebauh method
	HitungLuas() int
	HitungKeliling() int
}

// untuk memenuhi kontraknya kita bikin sebuah struct
type Persegi struct {
	Sisi int
}

// agar struct Persegi memenuhi aturan interface Luas, maka Persegi harus memiliki method HitungLuas()
func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

// misal kita bikin struct PersegiPanjang
type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {

}

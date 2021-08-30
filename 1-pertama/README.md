## Catatan

### Overview

- Setiap program go harus wajib masuk kedalam sebuah package
- main.go adalah main function untuk tempat mengeksekusi kode program golang
- go mod digunakan untuk menginisialisasi project golang
- caranya `go mod init <nama foldernya>`

### Run dan Build

- `go run <nama filenya>` yakni untuk menjalankan source code yang ada di file main.go
- `go build` perintah untuk membuatkan source code menjadi binary (.exe)
- untuk menjalankan hasil build `./<nama file buildnya>`

### Package

- package artinya paket
- fungsinya untuk mengelompokkan beberapa file yang mempunyai kemiripan menjadi dalam satu paket
- karena di `main.go` dan `entity.go` memiliki package main yang sama
- artinya `entity.go` bisa diakses didalam `main.go`
- asalkan dalam satu package bisa diakses langsung

### Running 2 package (package main + package lain)

- untuk merunning 2 package harus di running keduanya `go run main.go entity.go`
- di golang itu paling tidak ada 2 package

1. Package yang digunakan untuk Executable, artinya adalah program yang akan dieksekusi, yakni `package main`
2. Package yang sifatnya adalah Library, dimana bisa dipanggil dari file lain. Biasanya yang memanggil adalah source code dari package executable `selain main`

- Jadi semua package/program yang dieksekusi harus dimasukkan ke dalam package main

## Import

1. Untuk memanggil sebuah code standard libary bawaan golang, misal `fmt`
2. Untuk mengakses sebuah code tapi berada di package yang lain
3. Menggunakan library dari orang lain

## Func Main

- Yang secara default akan dijalankan saat kita merunning package executable
- Jadi, Func Main adalah method
- Saat kita menjalankan package main `go run main.go`
- Maka secara otomatis didalam package main akan dicari `Func Main`
- Lalu akan mengeksekusi kode program didalam func main
- Func Main inilah function yang utama untuk menaruh kode program aplikasi yang akan di execute, sifatnya wajib

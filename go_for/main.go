package main

import (
	"fmt"
)

func main() {
	// var keterangan string
	// keterangan := "Saya sedang belajar golang"

	// perulangan dengan for
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(keterangan, "=", i)
	// }

	// perulangan dengan while di golang
	// i := 1
	// for i <= 100 {
	// 	fmt.Println(keterangan, ":", i)
	// 	i++
	// }

	// perulangan dengan bentuk array
	title := "Golang the best language"

	// menggunakan variabel index dan letter
	for index, letter := range title {
		fmt.Println("index :", index, " letter :", string(letter))
	}

	// tidak menggunakan variabel index
	// tanda garis bawah, untuk menghilangkan variabel index atau tidak menggunakan variabel index hanya menampilkan variabel letter
	for _, letter := range title {
		fmt.Println("letter :", string(letter))
	}
}
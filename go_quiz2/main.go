package main

import (
	"fmt"
)

func main() {
	/**
	soal quiz 2 mengenai perulangan
	1. hanya mencetak huruf yang index nya berupa angka genap, sedangkan yang angka nya ganjil tidak di cetak?
	2. cetak hanya huruf vokal a,i, u, e, o, sedangkan selain huruf vokal tidak di cetak
	*/
	title := "Golang the best language"

	// for index, letter := range title {
		// untuk mengerjakan soal nomor 1 dengan menggunakan konsep modulo. dimana angka tersebut di bagi 2
		// if (index % 2 == 0) {
		// 	fmt.Println("index :", index, " letter :", string(letter))
		// }	
	// }
	// untuk mengerjakan nomor 2
	for index, letter := range title {
			letterString := string(letter)

			// menggunakan kondisi yang standar
			// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" || letterString == "" {
			// 	fmt.Println("index :", index, " letter :", string(letter))
			// }

			switch letterString {
			case "a","i","u","e","o","":
				fmt.Println("index :", index, " letter :", string(letter))
			}
		}

}